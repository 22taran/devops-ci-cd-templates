FROM php:8.2-cli-alpine AS build
WORKDIR /app
COPY --from=composer:latest /usr/bin/composer /usr/bin/composer
COPY composer.json composer.lock* ./
RUN composer install --no-dev --optimize-autoloader

FROM php:8.2-cli-alpine
WORKDIR /app
COPY --from=build /app/vendor ./vendor
COPY src/ ./src/
EXPOSE 8080
USER nobody
CMD ["php", "-S", "0.0.0.0:8080", "-t", "src/"]
