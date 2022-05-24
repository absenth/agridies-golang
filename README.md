# agridies - ARRL Field Day Logging Software
_____Agridies from the latin root AGRI - field and DIES - day.  Thus the name
of this project is literally "Field Day"_____

## Requirements

### Contest Requirements
1. Logs must be submitted in [Cabrillo Format](http://www.arrl.org/cabrillo-format-tutorial)
   **** [Example Cabrillo Log](https://a2a53e2b-2285-4083-9cff-c99fe5ba1658.filesusr.com/ugd/1c7085_6e6ab52ed6a246558704199c09aaf9f7.pdf)
2. Interface should only accept fields we're logging
   * BAND (ie. 20M)
   * MODE (ie. Phone)
   * DATE (2021-06-28)
   * TIME (0900)
   * SENT CALL (W1AW)
   * SENT CATEGORY (1A)
   * SENT SECTION (STX)
   * RECIEVED CALL (K9OOK)
   * RECIEVED CATEGORY  (1B)
   * RECIEVED SECTION (IN)
3. DATE & TIME should pull from the system, at the time we hit submit.
4. There Should be a Setup that lets us set one time Contest Data.

### Project Requirements
1. Power Loss of the Laptop should not result in losing logs.
2. Design for future multi-station logs (ie. one log, multiple computers)
3. Detect and Alert on Dupes (Same Callsign/Mode/Band)
   * look into `write ahead` or `live search`
4. It would be nice if we could pull BAND/Frequency data from the Rig.

### Research
1. let's use [fiber](https://gofiber.io/) as the http framework
   1. [Swagger](https://github.com/gofiber/swagger) framework for fiber.
2. Here's [another web framework](https://pkg.go.dev/github.com/gin-gonic/gin) called Gin
   1. [Swagger](https://github.com/swaggo/gin-swagger) framework for Gin. 
3. For Database connections [gorm](https://gorm.io/index.html)
4. Another option for Database connection [GoSQL](https://github.com/ilibs/gosql) seems like a better ORM
5. Here's a good [list of go packages](https://awesome-go.com/)
