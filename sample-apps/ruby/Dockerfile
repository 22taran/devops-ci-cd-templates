FROM ruby:3.3-alpine AS build
WORKDIR /app
COPY Gemfile Gemfile.lock* ./
RUN bundle install --without development test

FROM ruby:3.3-alpine
WORKDIR /app
COPY --from=build /usr/local/bundle /usr/local/bundle
COPY . .
EXPOSE 4567
USER nobody
CMD ["ruby", "app.rb", "-o", "0.0.0.0"]
