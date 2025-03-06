FROM maven:3.9.9-eclipse-temurin-21-alpine AS build
WORKDIR /app
COPY pom.xml pom.xml
COPY codegen codegen
COPY service service
RUN mvn clean package -DskipTests

FROM amazoncorretto:21-alpine
WORKDIR /app
COPY --from=build /app/service/target/*.jar app.jar
ENTRYPOINT ["java", "-jar", "app.jar"]
