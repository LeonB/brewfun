// var selector = '.std';
// var source = 'test';
// var defaultMatch = 0.5;
var root = document.querySelector(selector);
var iter = document.createNodeIterator (root, NodeFilter.SHOW_ELEMENT | NodeFilter.SHOW_TEXT);

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

var els = [];
while (child = iter.nextNode()) {
// [...root.children].forEach(function(child) {
    if (child.nodeName == "H3") {
        els[els.length] = [child];
        continue;
    }

    if (els.length > 0) {
        els[els.length-1].push(child);
    }
};

var substitutes = els.map(function(children) {
    var hopName = children[0].innerText;
    var country = null;
    // var hop = children[0].querySelector('a').name;
    var substitutes = [];

    children.forEach(function(child) {
        if (child.innerHTML == undefined) {
            return;
        }

        if (child.innerHTML.includes('Substitutes:')) {
            substitutes = child.nextSibling.textContent.split(',');
            substitutes = substitutes.map(function(sub) { return sub.trim(); });
            return;
        }
    });

    // skip hops without substitutes
    if (substitutes.length == 0) {
        return;
    }

    // try to get country from hop name
    var [hopName, country] = findCountry(hopName, country);

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
                country: null
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
