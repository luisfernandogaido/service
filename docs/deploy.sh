#!/usr/bin/env bash
file=/var/www/html/service.gaido.net.br/deploy/service
if [ -e "$file" ]; then
    systemctl stop service.service
    mv "$file" /var/www/html/service.gaido.net.br
    chmod 0774 /var/www/html/service.gaido.net.br/service
    systemctl start service.service
fi