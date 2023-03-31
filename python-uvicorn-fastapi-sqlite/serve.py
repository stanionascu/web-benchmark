import sqlite3

from fastapi import FastAPI, Request
from fastapi.responses import HTMLResponse
from fastapi.templating import Jinja2Templates

app = FastAPI()
templates = Jinja2Templates(directory="templates")
conn = sqlite3.connect("../db.sqlite3")
cursor = conn.cursor()

@app.get("/", response_class=HTMLResponse)
async def index(request: Request):
  films = cursor.execute("SELECT title,release_year FROM film LIMIT 100").fetchall()
  return templates.TemplateResponse("index.html", {"request": request, "name": request.query_params.get("name"), "films": films})
