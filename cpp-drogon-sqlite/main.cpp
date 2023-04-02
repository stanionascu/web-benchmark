#include <drogon/drogon.h>
#include <sqlite3.h>

#include <iostream>
#include <string>
#include <vector>

#include "Film.h"

using namespace drogon;

int main(int argc, char **argv) {
  sqlite3 *db = nullptr;

  if (sqlite3_open("../../db.sqlite3", &db) > 0) {
    std::cerr << sqlite3_errmsg(db) << std::endl;
    return 1;
  }

  app().registerHandler(
      "/?name={name}",
      [&](const HttpRequestPtr &req,
          std::function<void(const HttpResponsePtr &)> &&callback,
          const std::string &name) {
        std::vector<Film> films;
        films.reserve(100);
        sqlite3_exec(
            db, "SELECT title,release_year FROM film LIMIT 100",
            [](void *data, int argc, char **argv, char **) -> int {
              auto films = (std::vector<Film> *)data;
              films->emplace_back(Film{.index = static_cast<int>(films->size()),
                                       .title = argv[0],
                                       .releaseYear = argv[1]});
              return 0;
            },
            &films, nullptr);

        HttpViewData data;
        data.insert("name", name);
        data.insert("films", films);
        auto resp = HttpResponse::newHttpViewResponse("IndexHtml", data);
        callback(resp);
      });

  app()
      .setLogLevel(trantor::Logger::kWarn)
      .addListener("0.0.0.0", 3000)
      .setThreadNum(std::thread::hardware_concurrency())
      .run();

  sqlite3_close(db);
  return 0;
}
