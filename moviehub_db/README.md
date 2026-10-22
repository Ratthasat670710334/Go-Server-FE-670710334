# docker compose up -d

# sudo apt-get update && sudo apt-get install -y postgresql-client



psql "postgres://postgres:secret@127.0.0.1:5432/moviehub" -c "SELECT * FROM members;"