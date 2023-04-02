cmake_minimum_required(VERSION 3.11)
project(cpp_drogon_sqlite CXX)
set(CMAKE_CXX_STANDARD 20)

option(BUILD_REDIS OFF)
option(BUILD_POSTGRESQL OFF)
option(BUILD_ORM OFF)
option(BUILD_MYSQL OFF)
option(BUILD_SQLITE OFF)
option(BUILD_EXAMPLES "" OFF)
option(BUILD_DOC OFF)

include(FetchContent)
FetchContent_Declare(drogon
  GIT_REPOSITORY https://github.com/drogonframework/drogon.git
  GIT_TAG v1.8.4
)

find_package(ZLIB REQUIRED)
find_package(SQLite3 REQUIRED)
# find_package(Drogon REQUIRED PATHS ${CMAKE_CURRENT_BINARY_DIR}/drogon)
FetchContent_MakeAvailable(drogon)

add_executable(${PROJECT_NAME} main.cpp IndexHtml.cc IndexHtml.h)
target_link_libraries(${PROJECT_NAME} PUBLIC SQLite::SQLite3 ZLIB::ZLIB drogon)
