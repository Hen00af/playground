import "./style.css";

const CHARACTERS = [" ", ".", ":", "-", "=", "+", "*", "#", "%", "@"];

/*
 * 1文字あたり、縦横それぞれ何ピクセル調べるか。
 * 4なら 4 × 4 = 16ピクセルを平均する。
 */
const SUPERSAMPLE = 4;

const fileInput = document.querySelector<HTMLInputElement>("#file-input");
const widthInput = document.querySelector<HTMLInputElement>("#width-input");
const widthValue = document.querySelector<HTMLSpanElement>("#width-value");
const canvas = document.querySelector<HTMLCanvasElement>("#canvas");
const output = document.querySelector<HTMLPreElement>("#output");

if (!fileInput || !widthInput || !widthValue || !canvas || !output) {
  throw new Error("Required HTML element was not found.");
}

const context = canvas.getContext("2d", {
  willReadFrequently: true,
});

if (!context) {
  throw new Error("Canvas 2D context could not be created.");
}

let currentImage: HTMLImageElement | null = null;

const getBrightness = (
  red: number,
  green: number,
  blue: number,
): number => {
  return (
    (0.2126 * red +
      0.7152 * green +
      0.0722 * blue) /
    255
  );
};

const getCharacter = (brightness: number): string => {
  const index = Math.floor(
    brightness * (CHARACTERS.length - 1),
  );

  return CHARACTERS[index];
};

const getAverageBrightness = (
  imageData: ImageData,
  cellX: number,
  cellY: number,
  sampleWidth: number,
): number => {
  let brightnessSum = 0;

  for (let offsetY = 0; offsetY < SUPERSAMPLE; offsetY++) {
    for (let offsetX = 0; offsetX < SUPERSAMPLE; offsetX++) {
      const pixelX = cellX * SUPERSAMPLE + offsetX;
      const pixelY = cellY * SUPERSAMPLE + offsetY;

      const pixelIndex = (pixelY * sampleWidth + pixelX) * 4;

      const red = imageData.data[pixelIndex];
      const green = imageData.data[pixelIndex + 1];
      const blue = imageData.data[pixelIndex + 2];

      brightnessSum += getBrightness(red, green, blue);
    }
  }

  return brightnessSum / (SUPERSAMPLE * SUPERSAMPLE);
};

const renderAscii = (image: HTMLImageElement): void => {
  const columns = Number(widthInput.value);

  /*
   * 等幅フォントは縦長に見えるため、高さを半分程度に補正する。
   */
  const rows = Math.max(
    1,
    Math.round(
      (image.height / image.width) * columns * 0.5,
    ),
  );

  /*
   * ASCII文字数よりも高解像度な画像をCanvasに描く。
   *
   * columns = 100
   * SUPERSAMPLE = 4
   * の場合、Canvasの横幅は400ピクセルになる。
   */
  const sampleWidth = columns * SUPERSAMPLE;
  const sampleHeight = rows * SUPERSAMPLE;

  canvas.width = sampleWidth;
  canvas.height = sampleHeight;

  context.drawImage(
    image,
    0,
    0,
    sampleWidth,
    sampleHeight,
  );

  const imageData = context.getImageData(
    0,
    0,
    sampleWidth,
    sampleHeight,
  );

  const lines: string[] = [];

  for (let y = 0; y < rows; y++) {
    let line = "";

    for (let x = 0; x < columns; x++) {
      const brightness = getAverageBrightness(
        imageData,
        x,
        y,
        sampleWidth,
      );

      line += getCharacter(brightness);
    }

    lines.push(line);
  }

  output.textContent = lines.join("\n");
};

const loadImage = (file: File): void => {
  const image = new Image();
  const objectUrl = URL.createObjectURL(file);

  image.onload = () => {
    currentImage = image;
    renderAscii(image);

    URL.revokeObjectURL(objectUrl);
  };

  image.src = objectUrl;
};

fileInput.addEventListener("change", () => {
  const [file] = Array.from(fileInput.files ?? []);

  if (!file) {
    return;
  }

  loadImage(file);
});

widthInput.addEventListener("input", () => {
  widthValue.textContent = widthInput.value;

  if (currentImage) {
    renderAscii(currentImage);
  }
});
