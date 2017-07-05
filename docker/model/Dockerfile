FROM mariadb:latest
RUN echo "Europe/Paris" > /etc/timezone
RUN dpkg-reconfigure -f noninteractive tzdata
ADD init.sql /docker-entrypoint-initdb.d/
ARG dbname
ARG dbuser
ARG dbpassword
ENV MYSQL_ROOT_PASSWORD=root
ENV MYSQL_DATABASE=$dbname
ENV MYSQL_USER=$dbuser
ENV MYSQL_PASSWORD=$dbpassword
