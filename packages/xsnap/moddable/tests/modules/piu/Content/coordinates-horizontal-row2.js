/*---
description: 
flags: [onlyStrict]
---*/

/*
From https://github.com/Moddable-OpenSource/moddable/blob/public/documentation/piu/piu.md#coordinates

When a content's container is a row object:
 * left and right coordinates are relative to their previous and next properties
   top and bottom coordinates are relative to their container
   If height, top, and bottom coordinates are all specified, the height will be overruled
*/

const fixedWidth = 100;
const fixedHeight = 100;

const fixedSizeContent1 = new Content(null, { left: 0, width: fixedWidth, top: 0, height: fixedHeight });
const fixedSizeContent2 = new Content(null, { left: 0, width: fixedWidth, top: 0, height: fixedHeight });
const relativeSizeContent = new Content(null, { left: 0, right: 0, top: 0, bottom: 0 });

const fixedSizeRow = new Row(null, {
    top: 0, height: 300, left: 0, width: 300,
    contents: [
        fixedSizeContent1,
        relativeSizeContent,
        fixedSizeContent2
    ]
});

new Application(null, { contents: [fixedSizeRow] });

assert.sameValue(fixedSizeContent1.width, fixedWidth, `fixedSizeContent1 width should be ${fixedWidth}`);
assert.sameValue(fixedSizeContent2.width, fixedWidth, `fixedSizeContent2 width should be ${fixedWidth}`);
let expected = fixedSizeRow.width - fixedSizeContent1.width - fixedSizeContent2.width;
assert.sameValue(relativeSizeContent.width, expected, `relativeSizeContent width should be ${expected}`);
assert.sameValue(fixedSizeContent1.x, 0, `fixedSizeContent1 x should be 0`);
expected = relativeSizeContent.width + relativeSizeContent.x;
assert.sameValue(fixedSizeContent2.x, expected, `fixedSizeContent2 x should be ${expected}`);
expected = fixedSizeContent1.width + fixedSizeContent1.x
assert.sameValue(relativeSizeContent.x, expected, `relativeSizeContent x should be ${expected}`);

// Change width of fixed size contents in row
// This should have an effect on width/x coordinates of other contents in row
fixedSizeContent1.width = 0;
assert.sameValue(fixedSizeContent1.width, 0, `fixedSizeContent1 width should be 0`);
assert.sameValue(fixedSizeContent2.width, fixedWidth, `fixedSizeContent2 width should be ${fixedWidth}`);
expected = fixedSizeRow.width - fixedSizeContent1.width - fixedSizeContent2.width;
assert.sameValue(relativeSizeContent.width, expected, `relativeSizeContent width should be ${expected}`);
assert.sameValue(fixedSizeContent1.x, 0, `fixedSizeContent1 x should be 0`);
expected = relativeSizeContent.width + relativeSizeContent.x;
assert.sameValue(fixedSizeContent2.x, expected, `fixedSizeContent2 x should be ${expected}`);
expected = fixedSizeContent1.width + fixedSizeContent1.x;
assert.sameValue(relativeSizeContent.x, expected, `relativeSizeContent x should be ${expected}`);

fixedSizeContent2.width = 0;
assert.sameValue(fixedSizeContent1.width, 0, `fixedSizeContent1 width should be 0`);
assert.sameValue(fixedSizeContent2.width, 0, `fixedSizeContent2 width should be 0`);
expected = fixedSizeRow.width - fixedSizeContent1.width - fixedSizeContent2.width;
assert.sameValue(relativeSizeContent.width, expected, `relativeSizeContent width should be ${expected}`);
assert.sameValue(fixedSizeContent1.x, 0, `fixedSizeContent1 x should be 0`);
expected = relativeSizeContent.width + relativeSizeContent.x;
assert.sameValue(fixedSizeContent2.x, expected, `fixedSizeContent2 x should be ${expected}`);
expected = fixedSizeContent1.width + fixedSizeContent1.x;
assert.sameValue(relativeSizeContent.x, expected, `relativeSizeContent x should be ${expected}`);

// Change width/left margin of row itself
// This should have an effect on width/x coordinates of contents in row
fixedSizeRow.width = fixedWidth;
assert.sameValue(fixedSizeContent1.width, 0, `fixedSizeContent1 width should be 0`);
assert.sameValue(fixedSizeContent2.width, 0, `fixedSizeContent2 width should be 0`);
assert.sameValue(relativeSizeContent.width, fixedWidth, `relativeSizeContent width should be ${fixedWidth}`);
assert.sameValue(fixedSizeContent1.x, 0, `fixedSizeContent1 x should be 0`);
assert.sameValue(fixedSizeContent2.x, fixedWidth, `fixedSizeContent2 x should be ${fixedWidth}`);
assert.sameValue(relativeSizeContent.x, 0, `relativeSizeContent x should be 0`);

fixedSizeRow.moveBy(100, 0);
assert.sameValue(fixedSizeContent1.width, 0, `fixedSizeContent1 width should be 0`);
assert.sameValue(fixedSizeContent2.width, 0, `fixedSizeContent2 width should be 0`);
assert.sameValue(relativeSizeContent.width, fixedWidth, `relativeSizeContent width should be ${fixedWidth}`);
assert.sameValue(fixedSizeContent1.x, 100, `fixedSizeContent1 x should be 100`);
assert.sameValue(fixedSizeContent2.x, 200, `fixedSizeContent2 x should be 200`);
assert.sameValue(relativeSizeContent.x, 100, `relativeSizeContent x should be 100`);
