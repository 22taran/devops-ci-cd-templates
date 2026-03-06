require 'sinatra'

set :port, 4567

get '/' do
  content_type :json
  { message: greet('World') }.to_json
end

get '/health' do
  content_type :json
  { status: 'healthy' }.to_json
end

def greet(name = nil)
  return 'Hello, World!' if name.nil? || name.strip.empty?
  "Hello, #{name}!"
end

def add(a, b)
  a + b
end
