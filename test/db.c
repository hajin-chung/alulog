#include <stdio.h>
#include <sqlite3.h> 

static int callback(void *NotUsed, int argc, char **argv, char **azColName) {
  int i;
  for(i = 0; i<argc; i++) {
    printf("%s = %s\n", azColName[i], argv[i] ? argv[i] : "NULL");
  }
  printf("\n");
  return 0;
}

int main(int argc, char* argv[]) {
  sqlite3 *db;
  char *zErrMsg = 0;
  char *sql;
  int rc;

  rc = sqlite3_open("test.db", &db);
  if (rc != SQLITE_OK) {
    fprintf(stderr, "Can't open database: %s\n", sqlite3_errmsg(db));
    return 0;
  } else {
    fprintf(stderr, "Opened database successfully\n");
  }

  sql = "CREATE TABLE test("\
    "id  INT PRIMARY KEY NOT NULL,"\
    "title    TEXT       NOT NULL,"\
    "subtitle TEXT       NOT NULL,"\
    "content  TEXT       NOT NULL);";

  rc = sqlite3_exec(db, sql, callback, 0, &zErrMsg);
  if (rc != SQLITE_OK) {
    fprintf(stderr, "SQL Error: %s\n", zErrMsg);
    sqlite3_free(zErrMsg);
  } else {
    fprintf(stderr, "SQL success\n");
  }

  sqlite3_close(db);
  return 0;
}
