var defaultMatch = 0.5;
var source = 'https://www.homebrewersassociation.org/how-to-brew/hop-substitutions/';
var rows = document.querySelectorAll('table tr');

// select all rows but the first
var substitutes = [...rows].slice(1).map(function(row) {
    var hopName = row.querySelector('td:nth-child(1)').innerHTML;
    var substitutes = [...row.querySelectorAll('td:nth-child(2) li')].map(function(sub) { return sub.innerHTML; });

    // strip tags
    hopName = hopName.replace(/(<([^>]+)>)/ig,"");
    substitutes = substitutes.map(function(sub) { return sub.replace(/(<([^>]+)>)/ig,""); });

    // remove whitespace
    hopName = hopName.trim();
    substitutes = substitutes.map(function(sub) { return sub.trim(); });

    // match country
    var country = null;
    if (hopName.includes('(') && hopName.includes(')')) {
        var matches = hopName.match(/(.*)\s\((.*)\)/);
        hopName = matches[1];
        country = matches[2];
    }

    // convert substitutes to objects
    substitutes = substitutes.map(function(sub) {
        var name = sub;
        var match = defaultMatch;
        var country = null;

        // match country
        var country = null;
        if (name.includes('(') && name.includes(')')) {
            var matches = name.match(/(.*)\s\((.*)\)/);
            name = matches[1];
            country = matches[2];
        }

        return {
            hop: {
                name: name,
                country: country
            },
            match: match
        };
    });

    // return complete substitution object
    return {
        hop: {
            name: hopName,
            country: country,
        },
        substitutes: substitutes,
        source: source
    }
}).filter(function(sub) { return sub; });
