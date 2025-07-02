#!/bin/bash
echo "İndiriliyor..."

FILE_URL="https://github.com/isa-programmer/go-projects/raw/refs/heads/main/projects/pwish/pwish-go"

wget $FILE_URL
chmod +x pwish-go

echo "İndirildi! Çalıştırmak için ./pwish-go"
