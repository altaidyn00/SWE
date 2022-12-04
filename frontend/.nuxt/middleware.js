const middleware = {}

middleware['admin'] = require('../middleware/admin.js')
middleware['admin'] = middleware['admin'].default || middleware['admin']

middleware['loggedin'] = require('../middleware/loggedin.js')
middleware['loggedin'] = middleware['loggedin'].default || middleware['loggedin']

middleware['notadmin'] = require('../middleware/notadmin.js')
middleware['notadmin'] = middleware['notadmin'].default || middleware['notadmin']

middleware['notloggedin'] = require('../middleware/notloggedin.js')
middleware['notloggedin'] = middleware['notloggedin'].default || middleware['notloggedin']

export default middleware
