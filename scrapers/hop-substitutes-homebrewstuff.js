var defaultMatch = 0.5;
var root = document.querySelector(selector);
var iter = document.createNodeIterator (root, NodeFilter.SHOW_ELEMENT | NodeFilter.SHOW_TEXT);

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
    var hop = children[0].innerText;
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
            name: hop,
            country: null,
        },
        substitutes: substitutes,
        source: source
    };
}).filter(function(sub) { return sub; });
substitutes;
