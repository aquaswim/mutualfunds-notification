#!/bin/sh

echo "Set the crontab"
echo "${REPORT_CRONTAB:-0 17 * * *} /usr/local/bin/python /app/main.py" > /var/spool/cron/crontabs/root

echo "run cron"
crond -f