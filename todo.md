* Create test data with images and text
* Create basic title screen with simple ascii art
* Create UI that collects AWS credentials and saves locally. Runs once on startup and collects data
    * Access Key
    * Secret key
* Create basic cache that can reference local config & yaml
* Push test data to S3
* Pull data from S3 if missing, save locally
* Create UI for updating and pulling S3 data to local cache
* Display and update timestamp whenever cache is updated from S3
* View test data as tree
* Implement "Guess which event" that displays picture and has user guess from a list of options