# Install on Ubuntu

14.04 => trusty
16.04 => xenial
17.04 => zesty

1. Write below to /etc/apt/sources.list.d/pgdg.list:

deb http://apt.postgresql.org/pub/repos/apt/ zesty-pgdg main

2. wget --quiet -O - https://www.postgresql.org/media/keys/ACCC4CF8.asc | \
  sudo apt-key add -

3. sudo apt-get update

4. sudo apt install postgresql-9.6
