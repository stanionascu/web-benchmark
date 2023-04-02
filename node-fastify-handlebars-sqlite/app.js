'use strict'

const os = require('os');
const cluster = require('node:cluster');
const process = require('node:process');
const numCPUs = os.availableParallelism();

const fastify = require('fastify')

const fs = require('fs')
const sqlite3 = require('sqlite3')
const db = new sqlite3.Database("../db.sqlite3")
const handlebars = require('handlebars');
const index = handlebars.compile(fs.readFileSync('templates/index.html').toString('utf-8'))

if (cluster.isPrimary) {
  console.log(`Primary ${process.pid} is running`);

  // Fork workers.
  for (let i = 0; i < numCPUs; i++) {
    cluster.fork();
  }

  cluster.on('exit', (worker, code, signal) => {
    console.log(`worker ${worker.process.pid} died`);
  });
} else {
  const app = fastify({logger: false})
  app.register(require('@fastify/compress'), { global: true, encodings: ["gzip"], threshold: 0 })
  app.register(function(router, opts, done) {
    router.get('/', async function (request, reply) {
      const films = new Promise((resolve, reject) => {
        db.all("SELECT title,release_year FROM film LIMIT 100", (err, rows) => {
          if (err != null) {
            reject(err)
          }
          const films = rows.map((row) => {
            return {title: row.title, releaseYear: row.release_year}
          })
          resolve(films)
        })
      });
      return reply.type('text/html').send(index({ name: request.query['name'], films: await films }))
    })
    done()
  })

  app.ready(() => app.hasPlugin('@fastify/compress'))

  app.listen({ port: 3000, host: "0.0.0.0" }, function (err, address) {
    if (err) {
      app.log.error(err)
      process.exit(1)
    }
    console.log(`Worker ${process.pid} started listening on ${address}`);
  })
}
