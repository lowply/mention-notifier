package main

type Response interface {
	query(url string) error
}
