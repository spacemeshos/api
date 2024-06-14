const SwaggerUI = require('swagger-ui');
require('swagger-ui/dist/swagger-ui.css');
const spec = require('./api.swagger.json');

SwaggerUI({
    spec,
    dom_id: '#swagger',
});