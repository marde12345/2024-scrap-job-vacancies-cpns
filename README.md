# Job Vacancies Scrapper For CPNS 2024

## Overview

This project is a scrapper tool designed to collect job vacancies that match your degree and formation for CPNS 2024.

## Prerequisites

- Go 1.19

## Setup

1. **Clone the Repository**

    ```bash
    git clone git@github.com:marde12345/2024-scrap-job-vacancies-cpns.git
    ```

2. **Install Dependencies**

    Navigate to the project directory and run:

    ```bash
    go mod vendor
    ```

3. **Search for Your Targeted Degree and Formation**

    Update the script to search for job vacancies that match your degree and formation.

4. **Set Your Session and Cookies**

    Ensure your session and cookies are set properly within the script to access the necessary data.

5. **Run the Scrapper**

    Execute the script to scrape all eligible job vacancies for you.

6. **Result**

    The result will be created on `output.csv`

## Usage

1. **Customize Your Search:**

   Modify the search parameters in the script according to your targeted degree and formation.

2. **Run the Scraper:**

    ```bash
    go run main.go
    ```

3. **Review the Results:**

   The eligible job vacancies will be scraped and stored as specified in the script.

---

Happy Scraping!
