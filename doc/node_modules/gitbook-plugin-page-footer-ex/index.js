var footer = require('./assets/lib/plugin');
var moment = require('moment');

module.exports = {
    book: {
        assets: "./assets",
        css: ["style/plugin.css"]
    },
    hooks: {
        'page:before': function (page) {
            return footer(this, page);
        }
    },
    filters: {
        dateFormat: function (d, format) {
            return moment(d).format(format);
        }
    }
};