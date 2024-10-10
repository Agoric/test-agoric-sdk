/*---
description: 
flags: [onlyStrict]
---*/

/*
From https://github.com/Moddable-OpenSource/moddable/blob/public/documentation/piu/piu.md#coordinates

When a content's container is a column object:
 * top and bottom coordinates are relative to their previous and next properties
   left and right coordinates are relative to their container
   If width, left, and right coordinates are all specified, the width will be overruled
*/

const fixedWidth = 100;
const fixedHeight = 100;

const fixedSizeContent1 = new Content(null, { left: 0, width: fixedWidth, top: 0, height: fixedHeight });
const fixedSizeContent2 = new Content(null, { left: 0, width: fixedWidth, top: 0, height: fixedHeight });
const relativeSizeContent = new Content(null, { left: 0, right: 0, top: 0, bottom: 0 });

const absoluteSizeCol = new Column(null, {
    top: 0, height: 300, left: 0, width: 300,
    contents: [
        fixedSizeContent1,
        relativeSizeContent,
        fixedSizeContent2
    ]
});

new Application(null, { contents: [absoluteSizeCol] });

assert.sameValue(fixedSizeContent1.height, fixedHeight, `fixedSizeContent1 height should be ${fixedHeight}`);
assert.sameValue(fixedSizeContent2.height, fixedHeight, `fixedSizeContent2 height should be ${fixedHeight}`);
let expected = absoluteSizeCol.height - fixedSizeContent1.height - fixedSizeContent2.height;
assert.sameValue(relativeSizeContent.height, expected, `relativeSizeContent height should be ${expected}`);
assert.sameValue(fixedSizeContent1.y, 0, `fixedSizeContent1 y should be 0`);
expected = relativeSizeContent.height + relativeSizeContent.y;
assert.sameValue(fixedSizeContent2.y, expected, `fixedSizeContent2 y should be ${expected}`);
expected = fixedSizeContent1.height + fixedSizeContent1.y;
assert.sameValue(relativeSizeContent.y, expected, `relativeSizeContent y should be ${expected}`);

// Change height of fixed size contents in column
// This should have an effect on height/y coordinates of other contents in column
fixedSizeContent1.height = 0;
assert.sameValue(fixedSizeContent1.height, 0, `fixedSizeContent1 height should be 0`);
assert.sameValue(fixedSizeContent2.height, fixedHeight, `fixedSizeContent2 height should be ${fixedHeight}`);
expected = absoluteSizeCol.height - fixedSizeContent1.height - fixedSizeContent2.height;
assert.sameValue(relativeSizeContent.height, expected, `relativeSizeContent height should be ${expected}`);
assert.sameValue(fixedSizeContent1.y, 0, `fixedSizeContent1 y should be 0`);
expected = relativeSizeContent.height + relativeSizeContent.y;
assert.sameValue(fixedSizeContent2.y, expected, `fixedSizeContent2 y should be ${expected}`);
expected = fixedSizeContent1.height + fixedSizeContent1.y;
assert.sameValue(relativeSizeContent.y, expected, `relativeSizeContent y should be ${expected}`);

fixedSizeContent2.height = 0;
assert.sameValue(fixedSizeContent1.height, 0, `fixedSizeContent1 height should be 0`);
assert.sameValue(fixedSizeContent2.height, 0, `fixedSizeContent2 height should be 0`);
expected = absoluteSizeCol.height - fixedSizeContent1.height - fixedSizeContent2.height;
assert.sameValue(relativeSizeContent.height, expected, `relativeSizeContent height should be ${expected}`);
assert.sameValue(fixedSizeContent1.y, 0, `fixedSizeContent1 y should be 0`);
expected = relativeSizeContent.height + relativeSizeContent.y;
assert.sameValue(fixedSizeContent2.y, expected, `fixedSizeContent2 y should be ${expected}`);
expected = fixedSizeContent1.height + fixedSizeContent1.y;
assert.sameValue(relativeSizeContent.y, expected, `relativeSizeContent y should be ${expected}`);

// Change height/top margin of column itself
// This should have an effect on height/y coordinates of contents in column
absoluteSizeCol.height = fixedHeight;
assert.sameValue(fixedSizeContent1.height, 0, `fixedSizeContent1 height should be 0`);
assert.sameValue(fixedSizeContent2.height, 0, `fixedSizeContent2 height should be 0`);
assert.sameValue(relativeSizeContent.height, fixedHeight, `relativeSizeContent height should be ${fixedHeight}`);
assert.sameValue(fixedSizeContent1.y, 0, `fixedSizeContent1 y should be 0`);
expected = relativeSizeContent.height + relativeSizeContent.y;
assert.sameValue(fixedSizeContent2.y, expected, `fixedSizeContent2 y should be ${expected}`);
assert.sameValue(relativeSizeContent.y, 0, `relativeSizeContent y should be 0`);

absoluteSizeCol.moveBy(0, 100);
assert.sameValue(fixedSizeContent1.height, 0, `fixedSizeContent1 height should be 0`);
assert.sameValue(fixedSizeContent2.height, 0, `fixedSizeContent2 height should be 0`);
assert.sameValue(relativeSizeContent.height, fixedHeight, `relativeSizeContent height should be ${fixedHeight}`);
assert.sameValue(fixedSizeContent1.y, 100, `fixedSizeContent1 y should be 100`);
expected = relativeSizeContent.height + relativeSizeContent.y;
assert.sameValue(fixedSizeContent2.y, expected, `fixedSizeContent2 y should be ${expected}`);
assert.sameValue(relativeSizeContent.y, 100, `relativeSizeContent y should be 100`);
