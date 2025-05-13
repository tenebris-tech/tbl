# TBL

This repository contains Go packages for working with Tugboat Logic (https://www.tugboatlogic.com/) and will be
extended as new APIs become available. PRs are welcome provided that your contribution is under the MIT license.

Copyright (c) 2021 Tenebris Technologies Inc.

For custom evidence collection software development please contact us via https://tenebris.com

Please see the LICENSE file for licence information.

For an example of how to use the evidence package, please refer to example/main.go

Please note that the example code uses an evidence package function to load http-headers.json.
This file can be downloaded from TubBoat Logic when the custom evidence integration is created, and contains
the required credentials. A separate endpoint (URL) is required for each evidence task, 
but a single set of credentials can be used. An example of the file is included as example/http-headers-example.json. 

**DO NOT** use your personal username and password, only use the credentials provided by TugBoat Logic for the custom integration.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
