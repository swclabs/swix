import asyncio
import logging
from core import GreeterServe

if __name__ == '__main__':
    greeter = GreeterServe()
    logging.basicConfig(level=logging.INFO)
    asyncio.get_event_loop().run_until_complete(greeter.serve())
