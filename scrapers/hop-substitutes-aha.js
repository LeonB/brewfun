// var defaultMatch = 0.5;
var rows = document.querySelectorAll(selector);

function findCountry(hopName, country) {
    // match country
    if (hopName.includes('(') && hopName.includes(')')) {
        var matches = hopName.match(/(.*)\s\((.*)\)/);

        // @TODO: check if what's between the brackets is really a country and
        // not some annotation

        hopName = matches[1];
        country = matches[2];
    }

    return [hopName, country];
}

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

    // try to get country from hop name
    var [hopName, country] = findCountry(hopName, country);

    // convert substitutes to objects
    substitutes = substitutes.map(function(sub) {
        var name = sub;
        var match = defaultMatch;
        var country = null;

        // try to get country from hop name
        var [name, country] = findCountry(name, country);

        return {
            hopA: {
                name: null,
                country: null,
            },
            hopB: {
                name: name,
                country: country
            },
            match: match
        };
    });

    // return complete hopSubsitutes array
    return substitutes.map(function(sub) {
        sub.hopA = {
            name: hopName,
            country: country,
        };
        sub.source = source;
        return sub;
    });
});

// flatten
substitutes = [].concat.apply([], substitutes);

// remove empty values
substitutes = substitutes.filter(function(sub) { return sub; });

// output to be marshalled
substitutes;
