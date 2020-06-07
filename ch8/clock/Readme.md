# clock
Spin up a time-zone specific clock to report the current time over a specified
port on your local machine.

**Note** you do not need to install Go to run this application!

This program uses environment variables to specify the timezone.

A list of allowed timezones can be found [here](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones).

## Usage
```
TZ=<TIMEZONE> ./clock --port=<PORT>
```

## Example
- Run the following command to start a US/Eastern clock on localhost:8010
    ```
    TZ=US/Eastern go run clock.go --port=8010
    ```

- To run multiple clocks, be sure to set a different port number for each 
clock, and run each instance of the application in the background.
    ```
    TZ=US/Eastern ./clock --port=8010 & \
    TZ=Asia/Tokyo ./clock --port=8020 & \
    TZ=Europe/London ./clock --port=8030 & \
    TZ=US/Pacific ./clock --port=8040 &
    ```

## Related applications
- [clockwall](https://github.com/ulricksennick/tgpl-exercises/ch8/clockwall): 
Connect to multiple clock sources to display a wall of clocks, or world clock.

