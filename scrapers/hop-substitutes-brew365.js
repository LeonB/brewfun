// var defaultMatch = 0.8;
var rows = document.querySelectorAll('table tr');

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
var substitutes = [...rows].slice(1).map(function(row, i) {
    var hopName = row.querySelector('td:nth-child(1)').innerHTML.replace(/(<([^>]+)>)/ig,"");;
    var substitutes = row.querySelector('td:nth-child(2)').innerHTML.split(',');

    // remove whitespace
    hopName = hopName.trim();
    substitutes = substitutes.map(function(sub) { return sub.trim(); });

    // try to get country from hop name
    var [hopName, country] = findCountry(hopName, country);

    // remove unknown subtitutes
    substitutes = substitutes.filter(function(substitute, i) {
        if (substitute == '?') {
            return false;
        }
        if (substitute.includes('not sure')) {
            return false;
        }
        return true;
    });

    // skip hops without substitutes
    if (substitutes.length == 0) {
        return;
    }

    // convert substitutes to objects
    substitutes = substitutes.map(function(sub) {
        var name = sub;
        var match = defaultMatch;
        var country = null;

        // if the hop name includes a '?', remove that and bump down the match
        if (sub.includes('?')) {
            match = defaultMatch/2;
            name = sub.replace(' (?)', '').replace('(?)', '');
        }

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

