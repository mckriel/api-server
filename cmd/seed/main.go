package main

import (
	"flag"
	"log"
)

func main() {
	var (
		mongodb_only = flag.Bool("mongodb", false, "Seed MongoDB only")
		mysql_only   = flag.Bool("mysql", false, "Seed MySQL only")
		all          = flag.Bool("all", false, "Seed all databases")
	)
	flag.Parse()

	if !*mongodb_only && !*mysql_only && !*all {
		*all = true
	}

	log.Println("starting database seeding...")

	if *mongodb_only || *all {
		log.Println("seeding mongodb...")
		err := seed_mongodb()
		if err != nil {
			log.Fatalf("MongoDB seeding failed: %v", err)
		}
		log.Println("mongodb seeding completed")
	}

	if *mysql_only || *all {
		log.Println("seeding mysql...")
		err := seed_mysql()
		if err != nil {
			log.Fatalf("MySQL seeding failed: %v", err)
		}
		log.Println("mysql seeding completed")
	}

	log.Println("database seeding completed successfully")
}