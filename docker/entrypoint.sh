su postgres sh -c "createuser test_user"
if [ $? -ne 0 ];then
   sleep 15
   su postgres sh -c "createuser test_user"
fi
su postgres sh -c "psql -c \"ALTER USER test_user PASSWORD 'test_passwd'\""
su postgres sh -c "psql -c \"CREATE DATABASE test\""
su postgres sh -c "psql -c \"GRANT ALL PRIVILEGES ON DATABASE test to test_user;;\""


cd /w3bstream/build && ./srv-applet-mgr migrate
cd /w3bstream/build && ./srv-applet-mgr &

sleep 3

cd /w3bstream/frontend && pnpm start