FROM scratch
#
#   Cue docker container build
#
ADD https://raw.githubusercontent.com/bagder/ca-bundle/master/ca-bundle.crt /etc/ssl/certs/ca-bundle.crt
ADD cans     /usr/local/bin/cans
ADD www/app  /var/www/
CMD ["cans"]
EXPOSE 8000