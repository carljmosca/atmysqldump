package main

import (
    "log"
	"os"
	"os/exec"
	"time"
    "github.com/mileusna/crontab"
)

func main() {

	const ENV_ATMYSQLDUMP_JOB = "ATMYSQLDUMP_JOB"
	const ENV_MYSQL_DATABASE = "MYSQL_DATABASE"
	const ENV_MYSQL_USERNAME = "MYSQL_USERNAME"
	const ENV_MYSQL_PASSWORD = "MYSQL_PASSWORD"
	const ENV_MYSQL_BACKUP_DIRECTORY = "MYSQL_BACKUP_DIRECTORY"
	const ENV_MYSQL_BACKUP_DESTINATION_HOST = "MYSQL_BACKUP_DESTINATION_HOST"
	const ENV_MYSQL_BACKUP_DESTINATION_DIRECTORY = "MYSQL_BACKUP_DESTINATION_DIRECTORY"

	// get required environment variables	
	atmysqldumpJob := getEnvironmentVariable(ENV_ATMYSQLDUMP_JOB, true)
	mysqlDatabase := getEnvironmentVariable(ENV_MYSQL_DATABASE, true)
	mysqlUsername := getEnvironmentVariable(ENV_MYSQL_USERNAME, true)
	mysqlPassword := getEnvironmentVariable(ENV_MYSQL_PASSWORD, true)
	mysqlBackupDirectory := getEnvironmentVariable(ENV_MYSQL_BACKUP_DIRECTORY, true)

	// get optional environment variables
	mysqlBackupDestinationHost := getEnvironmentVariable(ENV_MYSQL_BACKUP_DESTINATION_HOST, false)
	mysqlBackupDestinationDirectory := getEnvironmentVariable(ENV_MYSQL_BACKUP_DESTINATION_DIRECTORY, false)

	ctab := crontab.New() // create cron table

    // MustAddJob - will panic on wrong syntax or problems with function/arguments
	ctab.MustAddJob(atmysqldumpJob, doBackup, mysqlDatabase, mysqlUsername, 
		mysqlPassword, mysqlBackupDirectory, mysqlBackupDestinationHost, mysqlBackupDestinationDirectory) 

	for (true) {

	}
}

func getEnvironmentVariable(key string, required bool) (string) {
	value, found := os.LookupEnv(key)
	if !found {
		log.Println(key + " is not set")
		if required {
			os.Exit(1)
		}
		return ""
	}
	return value
}

func doBackup(mysqlDatabase string, mysqlUsername string, mysqlPassword string, 
	mysqlBackupDirectory string, mysqlBackupDestinationHost string, 
	mysqlBackupDestinationDirectory string) (string) {

		log.Println("beginning backup for " + mysqlDatabase + " to " + mysqlBackupDirectory)
		cmd := exec.Command("mysqldump", mysqlDatabase, "-u", mysqlUsername, "-p" + mysqlPassword)

		current := time.Now()
		outputFileName := mysqlBackupDirectory + "/" + mysqlDatabase + "_" + 
			current.Format("20060102-150405") + ".dump"

		outfile, err := os.Create(outputFileName)
    	if err != nil {
        	panic(err)
    	}	
    	defer outfile.Close()
		cmd.Stdout = outfile
		
		cmd.Run()
		log.Println("backup complete")
		return outputFileName

}
