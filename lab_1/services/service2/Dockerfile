#
# Build stage
#

FROM maven:3.8.2-jdk-8 as build

COPY src /home/app/src
COPY pom.xml /home/app

RUN mvn -f /home/app/pom.xml clean package

#
# Package stage
#
FROM tomcat:alpine

LABEL gmail="misterchuy21@gmail.com"

ARG JAR_FILE=home/app/target/WebApp.war

COPY --from=build ${JAR_FILE} /usr/local/tomcat/webapps/

EXPOSE 8080

CMD ["catalina.sh", "run"]