#!/bin/sh
#init pg
if [ ! -d "/var/lib/postgresql_data/data" ]; then
   echo "PG data does not exsit!"
   cp -r /var/lib/postgresql/data /var/lib/postgresql_data/
   #rm -rf /var/lib/postgresql/13
else
   echo "PG data exists!"
fi
#rm -f /etc/postgresql/13/main/postgresql.conf
#ln -s /w3bstream/build_image/conf/postgresql.conf /etc/postgresql/13/main/postgresql.conf 
chown -R postgres:postgres /var/lib/postgresql_data
chmod -R 700 /var/lib/postgresql_data/data
#Start postgres
#pg_ctl start -D /var/lib/postgresql_data/data -l /var/lib/postgresql/log.log
su - postgres -c "pg_ctl start -D /var/lib/postgresql/data -l /var/lib/postgresql/log.log"


#Start mosquitto
/etc/init.d/mosquitto start
#sleep 10

cd /w3bstream && ./srv-applet-mgr migrate
#cd /w3bstream/build && ./srv-applet-mgr init_admin &
cd /w3bstream && ./srv-applet-mgr &

sleep 3

cd /w3bstream/frontend && pnpm start
