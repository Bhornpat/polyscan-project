// Scans a file for "API_KEY"
use std::env;
use std::fs;

fn main() {
    let args: Vec<String> = env::args().collect();
    if args.len() < 2 {
        println!("No file provided");
        return;
    }
    let path = &args[1];
    let content = fs::read_to_string(path).unwrap_or_default();
    let count = content.matches("API_KEY").count();
    println!("{{\"api_keys_found\": {}}}", count);
}

