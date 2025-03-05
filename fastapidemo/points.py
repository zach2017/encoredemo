from fastapi import FastAPI
from fastapi.responses import JSONResponse
from fastapi.middleware.cors import CORSMiddleware

app = FastAPI()

# Configure CORS
app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],  # Allows all origins
    allow_credentials=True,
    allow_methods=["*"],  # Allows all methods
    allow_headers=["*"],  # Allows all headers
)

@app.get("/getPoints")
async def get_points():
    data = {
        "sports": [
            {
                "name": "Riverwalk Stadium",
                "lon": -86.3078,
                "lat": 32.3802
            },
            {
                "name": "Cramton Bowl",
                "lon": -86.2988,
                "lat": 32.3835
            }
        ],
        "music": [
            {
                "name": "Capitol Sounds",
                "lon": -86.3021,
                "lat": 32.3773
            },
            {
                "name": "Montgomery Performing Arts Centre",
                "lon": -86.3080,
                "lat": 32.3748
            }
        ]
    }
    return JSONResponse(content=data, media_type="application/json")