import { db } from "$lib/db";
import type { Film } from "../app";

export const prerender = false;

export async function load({ url }) {
  const name = url.searchParams.get('name') ?? 'unknown';

  const rows = new Promise<Film[]>((resolve, reject) => {
    db.prepare("SELECT title,release_year FROM film LIMIT 100")
      .all((err: any, rows: any[]) => {
        if (err !== null) {
          console.error(err)
          reject(err);
        }
        let data = rows.map((row) => {
          return { title: row.title, releaseYear: row.release_year };
        });
        resolve(data)
      });
  });

  return {
    name: name,
    films: await rows,
  };
}
