require 'rack'

class RackApplication
  def call(env)
    pp env
    request = Rack::Request.new(env)
    pp request

    response = if request.path_info == '/'
                 body = "#{request.request_method}: Hello #{request.params['name']}"
                 Rack::Response.new(body, 200, {'Content-Type' => 'text/plain'})
               elsif request.path_info == '/admin'
                 body = "#{request.request_method}: Hello Admin page #{request.params['name']}"
                 Rack::Response.new(body, 200, {'Content-Type' => 'text/plain'})
               else
                 Rack::Response.new('Not Found', 404, {'Content-Type' => 'text/plain'})
               end

    response.finish
  end
end

class AdminMiddleware < Rack::Auth::Basic
  def initialize(app)
    super
  end

  def call(env)
    req = Rack::Request.new(env)
    if req.path == '/admin'
      super
    else
      @app.call(env)
    end
  end

end

use Rack::ShowStatus
use AdminMiddleware do |username, password|
  username == password
end

run RackApplication.new
