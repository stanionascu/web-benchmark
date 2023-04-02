cmake_minimum_required(VERSION 3.0)
project(cpp_crow_sqlite CXX)

configure_file(templates/index.html templates/index.html COPYONLY)

find_package(Crow REQUIRED)
find_package(ZLIB REQUIRED)
find_package(SQLite3 REQUIRED)

set(CROW_FEATURES compression)


add_definitions(-DCROW_ENABLE_COMPRESSION)
add_executable(${PROJECT_NAME} main.cpp)
target_link_libraries(${PROJECT_NAME} PUBLIC Crow::Crow SQLite::SQLite3 ZLIB::ZLIB)
