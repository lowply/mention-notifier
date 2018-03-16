#!/usr/bin/env ruby

require 'erb'
require 'yaml'

CONFIG="./config.yml"

def usage(var)
  puts "#{var} is empty. Please update #{CONFIG}"
  exit 1
end

def generate(config)
  u = <<~"EOL"
    BUCKET: # S3 bucket to upload the zipped binary
    LOGIN: # Your GitHub user name
    GITHUB_TOKEN: # GitHub personal access token
    SLACK_ENDPOINT: # Slack API endpoint
    GITHUB_ENDPOINT: https://api.github.com/notifications # Optional 
    REASON: mention # Optional
    POLLING: true # Optional
  EOL
  IO.write(config, u)
end

def main
  if File.exist?(CONFIG)
    y = YAML.load_file(CONFIG)
    required = ["BUCKET", "LOGIN", "GITHUB_TOKEN", "SLACK_ENDPOINT"]
    required.each do |v|
      usage(v) if y[v].nil?
    end

    IO.write("cf.yml", ERB.new(IO.read("./cf.yml.erb"), nil, "%").result(binding))
  else
    generate(CONFIG)
    puts "Successfully generated #{CONFIG}"
  end
end

main()
