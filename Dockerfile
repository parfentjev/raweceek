FROM maven:3.9.9-amazoncorretto-21-alpine AS build

WORKDIR /build
COPY pom.xml pom.xml
COPY src src
COPY spec spec
RUN mvn clean package -q

FROM amazoncorretto:21-alpine
WORKDIR /release
COPY --from=build /build/target/*.jar app.jar
ENTRYPOINT ["java", "-jar", "app.jar"]
