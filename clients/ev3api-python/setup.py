"""
    EV3 API
"""
import os

from setuptools import setup, find_packages  # noqa: H301

NAME = "ev3api-python"
is_tag = os.environ.get("GITHUB_REF_TYPE") == "tag"
VERSION = os.environ.get("GITHUB_REF_NAME") if is_tag else "0.0.0-dev"
# To install the library, run the following
#
# python setup.py install
#
# prerequisite: setuptools
# http://pypi.python.org/pypi/setuptools

REQUIRES = [
  "grpcio ~= 1.51.0"
]

setup(
    name=NAME,
    version=VERSION,
    description="EV3 gRPC Python client",
    author="Matthew Machivenyika",
    author_email="accounts@machivenyika.ch",
    url="https://github.com/Bermos/EV3-gRPC",
    keywords=["gRPC", "EV3", "EV3 API"],
    python_requires=">=3.6",
    install_requires=REQUIRES,
    packages=find_packages(exclude=["test", "tests"]),
    include_package_data=True,
    license="",
    long_description="""\
    """
)
