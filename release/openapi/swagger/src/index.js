const SwaggerUI = require('swagger-ui');
import './themes/swagger-ui.css';
const spec = require('./api.swagger.json');
SwaggerUI({
    spec,
    dom_id: '#swagger',
}); 