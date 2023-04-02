#include <iostream>

#include <crow.h>
#include <sqlite3.h>

int main(int argc, char **argv) {
  crow::SimpleApp app;
  app.use_compression(crow::compression::GZIP);

  sqlite3 *db = nullptr;
  if (sqlite3_open("../../db.sqlite3", &db) > 0) {
    std::cerr << sqlite3_errmsg(db) << std::endl;
    return 1;
  }

  auto index = crow::mustache::load("index.html");

  CROW_ROUTE(app, "/")
  ([&](const crow::request &req) {
    crow::mustache::context ctx;
    ctx["name"] = req.url_params.get("name");
    std::vector<crow::mustache::context> films;
    films.reserve(100);
    sqlite3_exec(db, "SELECT title,release_year FROM film LIMIT 100",
      [](void *data, int argc, char **argv, char**) -> int {
        auto films = (std::vector<crow::mustache::context>*)data;
        crow::mustache::context film;
        film["index"] = films->size();
        film["title"] = argv[0];
        film["releaseYear"] = argv[1];
        films->emplace_back(std::move(film));
      return 0;
    }, &films, nullptr);
    ctx["films"] = std::move(films);
    return index.render(ctx); });

  app.loglevel(crow::LogLevel::Warning).multithreaded().bindaddr("0.0.0.0").port(3000).run();
  sqlite3_close(db);
  return 0;
}