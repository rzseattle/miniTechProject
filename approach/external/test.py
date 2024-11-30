from transformers import AutoImageProcessor, AutoModelForImageClassification
from PIL import Image
import torch
import os
import sys

# Załaduj model i procesor obrazu
model_name = "Dharma20/vit-base-fruits-360"
processor = AutoImageProcessor.from_pretrained(model_name)
model = AutoModelForImageClassification.from_pretrained(model_name)


print("Analyzing " + sys.argv[1])

def classify_images(image_paths):
    """
    Klasyfikacja wielu obrazów.
    
    :param image_paths: Lista ścieżek do obrazów
    :return: Lista wyników klasyfikacji
    """
    images = []
    
    # Wczytaj i przetwórz obrazy
    for path in image_paths:
        try:
            image = Image.open(path).convert("RGB")
            images.append(image)
        except Exception as e:
            print(f"Błąd wczytywania obrazu {path}: {e}")
            continue

    # Przetwarzanie wsadowe
    inputs = processor(images=images, return_tensors="pt")

    # Przewidywanie
    outputs = model(**inputs)
    logits = outputs.logits
    predicted_classes = torch.argmax(logits, dim=-1)

    # Mapowanie etykiet na klasy
    results = [
        (path, model.config.id2label[predicted_classes[i].item()])
        for i, path in enumerate(image_paths)
    ]
    return results


#image_folder = "./test"
#image_paths = [os.path.join(image_folder, f) for f in os.listdir(image_folder)]
#predictions = classify_images(image_paths)

## Wyświetlenie wyników
#for path, label in predictions:
    #print(f"Obraz: {path} -> Przewidziana klasa: {label}")

predictions = classify_images([sys.argv[1]])

for path, label in predictions:
    print(f"{path}-->{label}")
