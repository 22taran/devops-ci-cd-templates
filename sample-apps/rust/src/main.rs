use actix_web::{web, App, HttpResponse, HttpServer};
use serde::Serialize;

#[derive(Serialize)]
struct Message {
    message: String,
}

#[derive(Serialize)]
struct Health {
    status: String,
}

pub fn greet(name: &str) -> String {
    if name.trim().is_empty() {
        return String::from("Hello, World!");
    }
    format!("Hello, {}!", name)
}

pub fn add(a: i32, b: i32) -> i32 {
    a + b
}

async fn home() -> HttpResponse {
    HttpResponse::Ok().json(Message {
        message: greet("World"),
    })
}

async fn health() -> HttpResponse {
    HttpResponse::Ok().json(Health {
        status: String::from("healthy"),
    })
}

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    println!("Server running on port 8080");
    HttpServer::new(|| {
        App::new()
            .route("/", web::get().to(home))
            .route("/health", web::get().to(health))
    })
    .bind("0.0.0.0:8080")?
    .run()
    .await
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_greet_with_name() {
        assert_eq!(greet("DevOps"), "Hello, DevOps!");
    }

    #[test]
    fn test_greet_empty() {
        assert_eq!(greet(""), "Hello, World!");
    }

    #[test]
    fn test_add() {
        assert_eq!(add(2, 3), 5);
    }
}
