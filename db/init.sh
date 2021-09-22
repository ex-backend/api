set -e

mongo <<EOF
use admin
db.createUser({
	user: 'root',
	pwd: '$DB_PWD',
	roles: [{role: 'readWrite', db: 'mainDB'}]
})
db = new Mongo().getDB('mainDB')
db.createCollection('paper', {})

db.paper.insert([{ "type": "printer", "size": "A1"}])
EOF
