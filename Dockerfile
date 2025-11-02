FROM maven:3.9.11-amazoncorretto-25-alpine AS build

WORKDIR /build
COPY pom.xml pom.xml
COPY src src
COPY spec spec
RUN mvn clean package -ntp -q

FROM maven:3.9.11-amazoncorretto-25-alpine
WORKDIR /release
COPY --from=build /build/target/*.jar app.jar
ENTRYPOINT ["java", "-jar", "app.jar"]
