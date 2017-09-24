var defaultMatch = 0.8;
var rows = document.querySelectorAll('table tr');

// select all rows but the first
var substitutes = [...rows].slice(1).map(function(row, i) {
    var hopName = row.querySelector('td:nth-child(1)').innerHTML.replace(/(<([^>]+)>)/ig,"");;
    var substitutes = row.querySelector('td:nth-child(2)').innerHTML.split(',');

    // remove whitespace
    hopName = hopName.trim();
    substitutes = substitutes.map(function(sub) { return sub.trim(); });

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
            country: null,
        },
        substitutes: substitutes,
        source: source
    };
}).filter(function(sub) { return sub; });
substitutes;
