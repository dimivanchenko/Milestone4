FROM ubuntu:20.04
MAINTAINER diva

# SNTP & timezone
ENV TZ=Europe/Kiev
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

# Postgres
RUN apt-get -y update
RUN apt-get -y install postgresql postgresql-contrib sudo openjdk-11-jdk wget supervisor nano
RUN pg_ctlcluster 12 main start

USER postgres
RUN /etc/init.d/postgresql start && psql -c "ALTER USER postgres PASSWORD 'postgres';" && psql -c "CREATE DATABASE ss_demo_1;"

USER root
RUN echo "host  all     all     0.0.0.0/0       md5" >> /etc/postgresql/12/main/pg_hba.conf
RUN sed -i "s/#listen_addresses = 'localhost'/listen_addresses = '*'/g" /etc/postgresql/12/main/postgresql.conf
RUN mkdir /var/lib/postgresql/12/main/data && chown postgres:postgres /var/lib/postgresql/12/main/data
#RUN sed -i "s/data_directory = '\/var\/lib\/postgresql\/12\/main'/data_directory = '\/var\/lib\/postgresql\/12\/main\/data'/g" /etc/postgresql/12/main/postgresql.conf
#RUN chown postgres:postgres /etc/postgresql/12/main/postgresql.conf


USER postgres
RUN /etc/init.d/postgresql restart

COPY --chmod=777 ./postgres_service.sh /usr/local/bin/postgres_service.sh

# Tomcat
USER root
RUN mkdir /opt/tomcat
RUN wget https://www-eu.apache.org/dist/tomcat/tomcat-9/v9.0.62/bin/apache-tomcat-9.0.62.tar.gz -P /tmp
RUN tar -xf /tmp/apache-tomcat-9.0.62.tar.gz -C /opt/tomcat/
RUN ln -s /opt/tomcat/apache-tomcat-9.0.62 /opt/tomcat/latest
RUN chmod +x /opt/tomcat/latest/bin/*.sh


# Configuration supervisor daemon
COPY ./supervisord.conf /etc/supervisor/conf.d/supervisord.conf

# Configuration supervisord Web UI
COPY --chmod=644 ./supervisord_webui.conf /etc/supervisor/supervisord.conf

# Ports for Postgres, Tomcat, Supervisord Web interface
EXPOSE 5432 8080 9001

# Supervisor daemon
CMD ["/usr/bin/supervisord"]
