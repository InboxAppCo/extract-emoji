var fontkit = require('fontkit');
var fs = require('fs');
var rimraf = require('rimraf');
var mkdirp = require('mkdirp');
var sprintf = require("sprintf-js").vsprintf;

var sizes = [48, 64, 96, 160]
var baseFolder = __dirname + '/images';
var font = fontkit.openSync('/System/Library/Fonts/Apple Color Emoji.ttf');

function writeImage(glyph, size, outputFolder) {
  // extract an image for this glyph
  var image = glyph.getImageForSize(size);
  
  if (image) {
    var fileName = sprintf('%04x', [glyph.codePoints[0]])
    fs.writeFileSync(outputFolder + 'emoji_' + fileName + '.png', image.data);
  }
}

// delete the images directory and re-create it
rimraf.sync(baseFolder);
mkdirp.sync(baseFolder);

sizes.forEach(function (size) {
  currentFolder = sprintf('%s/%dx%d/', [baseFolder, size, size]);
  mkdirp.sync(currentFolder)

  // write longest characters first so regex works correctly
  // first, the flag glyphs
  var glyphs = font.glyphsForString('🇨🇳🇩🇪🇪🇸🇫🇷🇬🇧🇮🇹🇯🇵🇰🇷🇷🇺🇺🇸');
  glyphs.forEach(function(currentGlyph) {
    writeImage(currentGlyph, size, currentFolder);
  });


  // next, keycap characters
  '#0123456789'.split('').forEach(function(c) {
    var glyphs = font.glyphsForString(c + '\u20E3');
    if (glyphs.length == 1)
      writeImage(glyphs[0], size, currentFolder);
  });

  // and finally, write individual glyphs in the character set
  font.characterSet.forEach(function(codePoint) {
    writeImage(font.glyphForCodePoint(codePoint), size, currentFolder);
  });
})