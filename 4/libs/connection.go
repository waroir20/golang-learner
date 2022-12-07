package database

import (
	"database/sql"
	_ "github.com/lib/pq" // This is needed to include the postgres driver in the runtime
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

//---------------------------------------------------------------------------------------
//---------------------------DO NOT MODIFY THIS FILE-------------------------------------
//---------------------------------------------------------------------------------------

// Connection - Used to store connection information
type Connection struct {
	cfg Config
	db  *gorm.DB
}

// New - Creates new connection object
func New(cfg Config) *Connection {
	return &Connection{cfg: cfg}
}

// Open - Connects to the Database
func (c *Connection) Open() error {
	// Open DB & connect
	connectionString, err := c.cfg.ConnectionString()
	if err != nil {
		return err
	}
	con, err := sql.Open("postgres", connectionString)
	if err != nil {
		return err
	}

	var namingStrat schema.NamingStrategy
	if c.cfg.DefaultSchema != "" {
		namingStrat = schema.NamingStrategy{
			TablePrefix:   c.cfg.DefaultSchema + ".", // schema name
			SingularTable: false,
		}
	}
	c.db, err = gorm.Open(postgres.New(postgres.Config{Conn: con}), &gorm.Config{NamingStrategy: namingStrat})
	return err
}

// Active - Get active gorm.DB connection
func (c *Connection) Active() *gorm.DB {
	return c.db
}

// Disconnect - Disconnects database connection
func (c *Connection) Disconnect() error {
	sqlDB, err := c.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
