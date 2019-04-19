using System;
using System.Collections.Generic;
using System.Dynamic;
using System.Collections.Generic;
using System.Linq;
using System.Net.Http;
using System.Security.Cryptography;
using System.Text;
using System.Threading.Tasks;
using System.Web;
using Famis.Model;
using Newtonsoft.Json;
using JsonMap = System.Collections.Generic.Dictionary<string, object>;

namespace Famis
{
    public class Service : IService
    {
        private readonly Credentials _creds;
        private readonly HttpClient _client;
        private AuthResponse _authResponse;
        private readonly string _baseUri;

        private JsonSerializerSettings ignoreNullSetting = new JsonSerializerSettings() {
            NullValueHandling = NullValueHandling.Ignore
        };

        private Formatting noFormatting = Formatting.None;


        public Service(string serviceUrl, string username, string password) {
            _creds = new Credentials(username, password);
            _client = new HttpClient();
            _baseUri = serviceUrl;
        }

        private async Task AuthorizeIfNeeded() {
            if (_authResponse != null && !_authResponse.IsExpired) {
                return;
            }

            var jsonBody = JsonConvert.SerializeObject(_creds);
            var result = await _client.PostAsync(
                BuildUri("MobileWebServices/api/Login"),
                new StringContent(jsonBody, Encoding.UTF8, "application/json"));

            var responseBody = await result.Content.ReadAsStringAsync();
            var authResponse = JsonConvert.DeserializeObject<AuthResponse>(responseBody);
            if (!authResponse.Result) {
                throw new Exception($"Authorization Failed: {authResponse.Message}");
            }

            _authResponse = authResponse;
            _client.DefaultRequestHeaders.Add(
                "Authorization",
                $"{authResponse.Item.TokenType} {authResponse.Item.AccessToken}");
        }

        public async Task<object> GetRecord(string endpoint, string filter, string expand) {
            var resp = await Get<object>(endpoint, filter, expand);
            if (resp.PageResults.Count > 0) {
                return resp.PageResults[0];
            }

            return null;
        }

        public async Task<List<object>> GetRecords(
            string endpoint,
            string filter = null,
            string expand = null,
            int offset = 0,
            int limit = -1) {
            var response = await Get<object>(
                endpoint, filter, expand, offset, limit);
            return response.PageResults;
        }

        public async Task<PagedResponse<Company>> GetCompanies(
            string filter = null,
            string expand = null,
            int offset = 0,
            int limit = -1) {
            return await Get<Company>(
                "apis/360facility/v1/companies", filter, expand, offset,
                limit);
        }

        public async Task<Company> GetCompany(string id, string expand = null) {
            return await GetOne<Company>("apis/360facility/v1/companies", id, expand);
        }

        public async Task<PagedResponse<Property>> GetProperties(
            string filter = null,
            string expand = null,
            int offset = 0,
            int limit = -1) {
            return await Get<Property>(
                "apis/360facility/v1/properties", filter, expand, offset,
                limit);
        }

        public async Task<Property> GetProperty(string id, string expand = null) {
            return await GetOne<Property>("apis/360facility/v1/properties", id, expand);
        }

        public async Task<PagedResponse<Floor>> GetFloors(
            string filter = null,
            string expand = null,
            int offset = 0,
            int limit = -1) {
            return await Get<Floor>("apis/360facility/v1/floors", filter, expand, offset, limit);
        }

        public async Task<Floor> GetFloor(string id, string expand = null) {
            return await GetOne<Floor>("apis/360facility/v1/floors", id, expand);
        }

        public async Task<PagedResponse<Country>> GetCountries(
            string filter = null,
            string expand = null,
            int offset = 0,
            int limit = -1) {
            return await Get<Country>(
                "apis/360facility/v1/countries", filter, expand, offset,
                limit);
        }

        public async Task<Country> GetCountry(string id, string expand = null) {
            return await GetOne<Country>("apis/360facility/v1/countries", id, expand);
        }

        public async Task<PagedResponse<Crew>> GetCrews(
            string filter = null,
            string expand = null,
            int offset = 0,
            int limit = -1) {
            return await Get<Crew>("apis/360facility/v1/crews", filter, expand, offset, limit);
        }

        public async Task<Crew> GetCrew(string id, string expand = null) {
            return await GetOne<Crew>("apis/360facility/v1/crews", id, expand);
        }

        public async Task<PagedResponse<InvoiceHeader>> GetInvoiceHeaders(
            string filter = null,
            string expand = null,
            int offset = 0,
            int limit = -1) {
            return await Get<InvoiceHeader>(
                "apis/360facility/v1/invoiceheaders", filter, expand,
                offset, limit);
        }

        public async Task<InvoiceHeader> GetInvoiceHeader(string id, string expand = null) {
            return await GetOne<InvoiceHeader>("apis/360facility/v1/invoiceheaders", id, expand);
        }

        public async Task<PagedResponse<InvoiceLine>> GetInvoiceLines(
            string filter = null,
            string expand = null,
            int offset = 0,
            int limit = -1) {
            return await Get<InvoiceLine>(
                "apis/360facility/v1/invoicelines", filter, expand,
                offset, limit);
        }

        public async Task<InvoiceLine> GetInvoiceLine(string id, string expand = null) {
            return await GetOne<InvoiceLine>("apis/360facility/v1/invoicelines", id, expand);
        }

        public async Task<PagedResponse<JournalEntry>> GetJournalEntries(
            string filter = null,
            string expand = null,
            int offset = 0,
            int limit = -1) {
            return await Get<JournalEntry>(
                "apis/360facility/v1/journalentrynpfa", filter, expand,
                offset, limit);
        }

        public async Task<JournalEntry> GetJournalEntry(string id, string expand = null) {
            return await GetOne<JournalEntry>("apis/360facility/v1/journalentrynpfa", id, expand);
        }

        public async Task<PagedResponse<MaterialItem>> GetMaterialItems(
            string filter = null,
            string expand = null,
            int offset = 0,
            int limit = -1) {
            return await Get<MaterialItem>(
                "apis/360facility/v1/materialitems", filter, expand,
                offset, limit);
        }

        public async Task<MaterialItem> GetMaterialItem(string id, string expand = null) {
            return await GetOne<MaterialItem>("apis/360facility/v1/materialitems", id, expand);
        }

        public async Task<PagedResponse<PoHeader>> GetPoHeaders(
            string filter = null,
            string expand = null,
            int offset = 0,
            int limit = -1) {
            return await Get<PoHeader>(
                "apis/360facility/v1/purchaseorderheaders", filter, expand,
                offset, limit);
        }

        public async Task<PoHeader> GetPoHeader(string id, string expand = null) {
            return await GetOne<PoHeader>("apis/360facility/v1/purchaseorderheaders", id, expand);
        }

        public async Task<PagedResponse<PoLine>> GetPoLines(
            string filter = null,
            string expand = null,
            int offset = 0,
            int limit = -1) {
            return await Get<PoLine>(
                "apis/360facility/v1/purchaseorderlines", filter, expand,
                offset, limit);
        }

        public async Task<PoLine> GetPoLine(string id, string expand = null) {
            return await GetOne<PoLine>("apis/360facility/v1/purchaseorderlines", id, expand);
        }

        public async Task<PagedResponse<Receipt>> GetPoReceipts(
            string filter = null,
            string expand = null,
            int offset = 0,
            int limit = -1) {
            return await Get<Receipt>(
                "apis/360facility/v1/purchaseorderreceipts", filter, expand,
                offset,
                limit);
        }

        public async Task<Receipt> GetPoReceipt(string id, string expand = null) {
            return await GetOne<Receipt>("apis/360facility/v1/purchaseorderreceipts", id, expand);
        }

        public async Task<PagedResponse<PoStatus>> GetPoStatuses(
            string filter = null,
            string expand = null,
            int offset = 0,
            int limit = -1) {
            return await Get<PoStatus>(
                "apis/360facility/v1/purchaseorderstatuses", filter, expand,
                offset, limit);
        }

        public async Task<PoStatus> GetPoStatus(string id, string expand = null) {
            return await GetOne<PoStatus>("apis/360facility/v1/purchaseorderstatuses", id, expand);
        }

        public async Task<PagedResponse<PrHeader>> GetPrHeaders(
            string filter = null,
            string expand = null,
            int offset = 0,
            int limit = -1) {
            return await Get<PrHeader>(
                "apis/360facility/v1/purchaserequisitionheaders", filter,
                expand, offset,
                limit);
        }

        public async Task<PrHeader> GetPrHeader(string id, string expand = null) {
            return await GetOne<PrHeader>(
                "apis/360facility/v1/purchaserequisitionheaders", id,
                expand);
        }

        public async Task<PagedResponse<PrLine>> GetPrLines(
            string filter = null,
            string expand = null,
            int offset = 0,
            int limit = -1) {
            return await Get<PrLine>(
                "apis/360facility/v1/purchaserequisitionlines", filter, expand,
                offset,
                limit);
        }

        public async Task<PrLine> GetPrLine(string id, string expand = null) {
            return await GetOne<PrLine>("apis/360facility/v1/purchaserequisitionlines", id, expand);
        }

        public async Task<PagedResponse<State>> GetStates(
            string filter = null,
            string expand = null,
            int offset = 0,
            int limit = -1) {
            return await Get<State>("apis/360facility/v1/states", filter, expand, offset, limit);
        }

        public async Task<State> GetState(string id, string expand = null) {
            return await GetOne<State>("apis/360facility/v1/states", id, expand);
        }

        public async Task<PagedResponse<Uom>> GetUoms(
            string filter = null,
            string expand = null,
            int offset = 0,
            int limit = -1) {
            return await Get<Uom>("apis/360facility/v1/uoms", filter, expand, offset, limit);
        }

        public async Task<Uom> GetUom(string id, string expand = null) {
            return await GetOne<Uom>("apis/360facility/v1/uoms", id, expand);
        }

        public async Task<PagedResponse<AccountSegmentValueNPFA>> GetAcctSegs(
            string filter = null,
            string expand = null,
            int offset = 0,
            int limit = -1) {
            return await Get<AccountSegmentValueNPFA>(
                "apis/360facility/v1/accountsegmentvaluenpfa",
                filter, expand, offset, limit);
        }

        public async Task<AccountSegmentValueNPFA> GetAcctSeg(string id, string expand = null) {
            return await GetOne<AccountSegmentValueNPFA>(
                "apis/360facility/v1/accountsegmentvaluenpfa", id, expand);
        }

        public async Task<PagedResponse<User>> GetUsers(
            string filter = null,
            string expand = null,
            int offset = 0,
            int limit = -1) {
            return await Get<User>(
                "apis/360facility/v1/users", filter, expand, offset,
                limit);
        }

        public async Task<User> GetUser(string id, string expand = null) {
            return await GetOne<User>("apis/360facility/v1/users", id, expand);
        }

        public async Task<UpsertResponse<User>> PostUser(User user) {
            var body = JsonConvert.SerializeObject(user, noFormatting, ignoreNullSetting);
            await AuthorizeIfNeeded();
            var url = _baseUri + "/apis/360facility/v1/users?";
            StringContent content = new StringContent(body, Encoding.UTF8, "application/json");
            var response = await _client.PostAsync(url, content);
            if (response.IsSuccessStatusCode) {
                var userResponse =
                    JsonConvert.DeserializeObject<User>(await response.Content.ReadAsStringAsync());
                return new UpsertResponse<User>(true, null, userResponse);
            }

            var failResponse =
                JsonConvert.DeserializeObject<FamisHttpResponse>(
                    await response.Content.ReadAsStringAsync());
            return new UpsertResponse<User>(false, failResponse.Message, user);
        }

        public async Task<UpsertResponse<User>> PatchUser(User user, int id) {
            var body = JsonConvert.SerializeObject(user, noFormatting, ignoreNullSetting);
            await AuthorizeIfNeeded();
            var url = _baseUri + "/apis/360facility/v1/users?key=" + id;
            var method = new HttpMethod("PATCH");
            var request = new HttpRequestMessage(method, url) {
                Content = new StringContent(body, Encoding.UTF8, "application/json")
            };
            HttpResponseMessage response = await _client.SendAsync(request);
            if (response.IsSuccessStatusCode) {
                Console.WriteLine(response.Content.ToString());
                var successUser =
                    JsonConvert.DeserializeObject<User>(await response.Content.ReadAsStringAsync());
                return new UpsertResponse<User>(true, null, successUser);
            }

            var failResponse =
                JsonConvert.DeserializeObject<FamisHttpResponse>(
                    await response.Content.ReadAsStringAsync());
            return new UpsertResponse<User>(false, failResponse.Message, user);
        }

        public async Task<PagedResponse<WorkOrder>> GetWorkOrders(
            string filter = null,
            string expand = null,
            int offset = 0,
            int limit = -1) {
            return await Get<WorkOrder>(
                "apis/360facility/v1/workorders", filter, expand, offset,
                limit);
        }

        public async Task<WorkOrder> GetWorkOrder(string id, string expand = null) {
            return await GetOne<WorkOrder>("apis/360facility/v1/workorders", id, expand);
        }

        public async Task<PagedResponse<Warehouse>> GetWarehouses(
            string filter = null,
            string expand = null,
            int offset = 0,
            int limit = -1) {
            return await Get<Warehouse>(
                "apis/360facility/v1/warehouses", filter, expand, offset,
                limit);
        }

        public async Task<Warehouse> GetWarehouse(string id, string expand = null) {
            return await GetOne<Warehouse>("apis/360facility/v1/warehouses", id, expand);
        }

        public async Task<PagedResponse<Receipt>> GetReceipts(
            string filter = null,
            string expand = null,
            int offset = 0,
            int limit = -1) {
            return await Get<Receipt>(
                "apis/360facility/v1/purchaseorderreceipts", filter, expand,
                offset,
                limit);
        }

        public async Task<Receipt> GetReceipt(string id, string expand = null) {
            return await GetOne<Receipt>("apis/360facility/v1/purchaseorderreceipts", id, expand);
        }

        public async Task<PagedResponse<RequestType>> GetRequestTypes(
            string filter = null,
            string expand = null,
            int offset = 0,
            int limit = -1) {
            return await Get<RequestType>(
                "apis/360facility/v1/requesttypes", filter, expand,
                offset,
                limit);
        }

        public async Task<PagedResponse<OtherCostType>> GetOtherCostTypes(
            string filter = null,
            string expand = null,
            int offset = 0,
            int limit = -1) {
            return await Get<OtherCostType>(
                "apis/360facility/v1/othercosttypes", filter, expand,
                offset,
                limit);
        }

        public async Task<PagedResponse<RequestPriority>> GetRequestPriorities(
            string filter = null,
            string expand = null,
            int offset = 0,
            int limit = -1) {
            return await Get<RequestPriority>(
                "apis/360facility/v1/requestpriorities", filter,
                expand, offset,
                limit);
        }

        public async Task<PagedResponse<CrewUserAssociation>> GetCrewUserAssociations(
            string filter = null,
            string expand = null,
            int offset = 0,
            int limit = -1) {
            return await Get<CrewUserAssociation>(
                "apis/360facility/v1/crewuserassociations",
                filter, expand, offset,
                limit);
        }

        public async Task<PagedResponse<LaborEntry>> GetLaborEntries(
            string filter = null,
            string expand = null,
            int offset = 0,
            int limit = -1) {
            return await Get<LaborEntry>(
                "apis/360facility/v1/laborentries", filter, expand, offset,
                limit);
        }

        public async Task<PagedResponse<LaborReason>> GetLaborReasons(
            string filter = null,
            string expand = null,
            int offset = 0,
            int limit = -1) {
            return await Get<LaborReason>(
                "apis/360facility/v1/laborreasons", filter, expand,
                offset,
                limit);
        }

        public async Task<PagedResponse<AccountSegmentNPFA>> GetAcctSegNpfas(
            string filter = null,
            string expand = null,
            int offset = 0,
            int limit = -1) {
            return await Get<AccountSegmentNPFA>(
                "apis/360facility/v1/accountsegmentnpfa", filter,
                expand, offset, limit);
        }

        public async Task<AccountSegmentNPFA> GetAcctSegNpfa(string id, string expand = null) {
            return await GetOne<AccountSegmentNPFA>(
                "apis/360facility/v1/accountsegmentnpfa", id,
                expand);
        }

        private async Task<PagedResponse<T>> Get<T>(
            string endpoint,
            string filter = null,
            string expand = null,
            int offset = 0,
            int limit = -1) {
            await AuthorizeIfNeeded();

            var url = BuildUri(endpoint, filter, expand, offset, limit);
            Console.WriteLine(url);
            var body =
                await _client.GetStringAsync(url);


            var odataResp = JsonConvert.DeserializeObject<ODataResponse<T>>(body);
            Console.WriteLine(
                $"Received {odataResp.Value.Count} records. Next page {odataResp.NextLink}");
            return new PagedResponse<T>(odataResp.Value, odataResp.NextLinkUrl, this);
        }

        private async Task<T> GetOne<T>(
            string endpoint,
            string id,
            string expand = null,
            int offset = 0,
            int limit = 1000) {
            await AuthorizeIfNeeded();

            var body =
                await _client.GetStringAsync(
                    BuildUri(
                        endpoint, $"Id eq {id}", expand, offset,
                        limit));
            var odataResp = JsonConvert.DeserializeObject<ODataResponse<T>>(body);
            if (odataResp.Value.Count > 0) {
                return odataResp.Value.First();
            }

            return default(T);
        }

        public async Task<PagedResponse<T>> GetNext<T>(PagedResponse<T> current) {
            await AuthorizeIfNeeded();
            Console.WriteLine($"Requesting more from: {current.NextLink}");
            var body =
                await _client.GetStringAsync(current.NextLink);
            var odataResp = JsonConvert.DeserializeObject<ODataResponse<T>>(body);
            return new PagedResponse<T>(odataResp.Value, odataResp.NextLinkUrl, this);
        }

        public async Task<HttpResponseMessage> PostAsync(string entityType, string body) {
            await AuthorizeIfNeeded();
            var url = _baseUri + "/apis/360facility/v1/" + entityType.ToLower();
            StringContent content = new StringContent(body, Encoding.UTF8, "application/json");
            HttpResponseMessage response = await _client.PostAsync(url, content);
            return response;
        }

        private Uri BuildUri(
            string endpoint,
            string filter = null,
            string expand = null,
            int offset = 0,
            int limit = 1000) {
            var uri = new UriBuilder($"{_baseUri}/{endpoint}");
            var qParams = HttpUtility.ParseQueryString(string.Empty);
            if (!string.IsNullOrEmpty(filter)) {
                qParams["$filter"] = filter;
            }

            if (!string.IsNullOrEmpty(expand)) {
                qParams["$expand"] = expand;
            }

            qParams["$skip"] = offset.ToString();
            if (limit != -1) {
                qParams["$top"] = limit.ToString();
            }

            uri.Query = qParams.ToString();
            return uri.Uri;
        }

        public async Task<UpsertResponse<InvoiceHeader>> PostInvoice(InvoiceHeader famisInv) {
            var body = JsonConvert.SerializeObject(famisInv);
            await AuthorizeIfNeeded();
            var url = _baseUri + "/apis/360facility/v1/invoiceheaders?";
            StringContent content = new StringContent(body, Encoding.UTF8, "application/json");
            HttpResponseMessage response = await _client.PostAsync(url, content);
            if (response.IsSuccessStatusCode) {
                var successInv =
                    JsonConvert.DeserializeObject<InvoiceHeader>(
                        await response.Content.ReadAsStringAsync());
                return new UpsertResponse<InvoiceHeader>(true, null, successInv);
            }

            var failResponse =
                JsonConvert.DeserializeObject<FamisHttpResponse>(
                    await response.Content.ReadAsStringAsync());
            return new UpsertResponse<InvoiceHeader>(false, failResponse.Message, null);
        }

        public async Task<UpsertResponse<InvoiceLine>> PostInvoiceLine(InvoiceLine line) {
            var body = JsonConvert.SerializeObject(line);
            await AuthorizeIfNeeded();
            var url = _baseUri + "/apis/360facility/v1/invoicelines?";
            StringContent content = new StringContent(body, Encoding.UTF8, "application/json");
            HttpResponseMessage response = await _client.PostAsync(url, content);
            if (response.IsSuccessStatusCode) {
                var successInvLine =
                    JsonConvert.DeserializeObject<InvoiceLine>(
                        await response.Content.ReadAsStringAsync());
                return new UpsertResponse<InvoiceLine>(true, null, successInvLine);
            }

            var failResponse =
                JsonConvert.DeserializeObject<FamisHttpResponse>(
                    await response.Content.ReadAsStringAsync());
            return new UpsertResponse<InvoiceLine>(false, failResponse.Message, null);
        }

        public async Task<UpsertResponse<PoHeader>> PostPo(PoHeader famisPo) {
            var body = JsonConvert.SerializeObject(famisPo);
            await AuthorizeIfNeeded();
            var url = _baseUri + "/apis/360facility/v1/purchaseorderheaders?";
            var content = new StringContent(body, Encoding.UTF8, "application/json");
            var response = await _client.PostAsync(url, content);

            if (response.IsSuccessStatusCode) {
                var successPo =
                    JsonConvert.DeserializeObject<PoHeader>(
                        await response.Content.ReadAsStringAsync());
                return new UpsertResponse<PoHeader>(true, null, successPo);
            }

            var failResponse =
                JsonConvert.DeserializeObject<FamisHttpResponse>(
                    await response.Content.ReadAsStringAsync());
            return new UpsertResponse<PoHeader>(false, failResponse.Message, null);
        }

        public async Task<UpsertResponse<PoLine>> PostPoLine(POLineCreateRequest line) {
            var body = JsonConvert.SerializeObject(line);
            await AuthorizeIfNeeded();
            var url = _baseUri + "/apis/360facility/v1/purchaseorderlines?";
            StringContent content = new StringContent(body, Encoding.UTF8, "application/json");
            var response = await _client.PostAsync(url, content);
            if (response.IsSuccessStatusCode) {
                var successPoLine =
                    JsonConvert.DeserializeObject<PoLine>(
                        await response.Content.ReadAsStringAsync());
                return new UpsertResponse<PoLine>(true, null, successPoLine);
            }

            var failResponse =
                JsonConvert.DeserializeObject<FamisHttpResponse>(
                    await response.Content.ReadAsStringAsync());
            return new UpsertResponse<PoLine>(false, failResponse.Message, null);
        }

        public async Task<UpsertResponse<WorkOrder>> PatchProjectWriteBack(string json, int Id) {
            var body = json.ToString();
            await AuthorizeIfNeeded();
            var url = _baseUri + "/apis/360facility/v1/workorders?key=" + Id;
            var method = new HttpMethod("PATCH");

            var request = new HttpRequestMessage(method, url) {
                Content = new StringContent(body, Encoding.UTF8, "application/json")
            };
            HttpResponseMessage response = await _client.SendAsync(request);
            if (response.IsSuccessStatusCode) {
                var successWo =
                    JsonConvert.DeserializeObject<WorkOrder>(
                        await response.Content.ReadAsStringAsync());
                return new UpsertResponse<WorkOrder>(true, null, successWo);
            }

            var failResponse =
                JsonConvert.DeserializeObject<FamisHttpResponse>(
                    await response.Content.ReadAsStringAsync());
            return new UpsertResponse<WorkOrder>(false, failResponse.Message, null);
        }

        public async Task<UpsertResponse<WorkOrder>> PostProjectCW(ChildWorkOrder wo) {
            var body = JsonConvert.SerializeObject(wo);
            await AuthorizeIfNeeded();
            var url = _baseUri + "/apis/360facility/v1/workorders?";
            var content = new StringContent(body, Encoding.UTF8, "application/json");
            var response = await _client.PostAsync(url, content);

            if (response.IsSuccessStatusCode) {
                var successWo =
                    JsonConvert.DeserializeObject<WorkOrder>(
                        await response.Content.ReadAsStringAsync());
                return new UpsertResponse<WorkOrder>(true, null, successWo);
            }

            var failResponse =
                JsonConvert.DeserializeObject<FamisHttpResponse>(
                    await response.Content.ReadAsStringAsync());
            return new UpsertResponse<WorkOrder>(false, failResponse.Message, null);
        }

        async Task<UpsertResponse<Company>> IService.PatchCompany(Company company, int Id) {
            var body = JsonConvert.SerializeObject(company, noFormatting, ignoreNullSetting);
            await AuthorizeIfNeeded();
            var url = _baseUri + "/apis/360facility/v1/companies?key=" + Id;
            var method = new HttpMethod("PATCH");

            var request = new HttpRequestMessage(method, url) {
                Content = new StringContent(body, Encoding.UTF8, "application/json")
            };
            HttpResponseMessage response = await _client.SendAsync(request);
            if (response.IsSuccessStatusCode) {
                var successCom =
                    JsonConvert.DeserializeObject<Company>(
                        await response.Content.ReadAsStringAsync());
                return new UpsertResponse<Company>(true, null, successCom);
            }

            var failResponse =
                JsonConvert.DeserializeObject<FamisHttpResponse>(
                    await response.Content.ReadAsStringAsync());
            return new UpsertResponse<Company>(false, failResponse.Message, null);
        }

        public async Task<UpsertResponse<Company>> PostCompany(Company com) {
            var body = JsonConvert.SerializeObject(com);
            await AuthorizeIfNeeded();
            var url = _baseUri + "/apis/360facility/v1/companies?";
            var content = new StringContent(body, Encoding.UTF8, "application/json");
            var response = await _client.PostAsync(url, content);

            if (response.IsSuccessStatusCode) {
                var successCom =
                    JsonConvert.DeserializeObject<Company>(
                        await response.Content.ReadAsStringAsync());
                return new UpsertResponse<Company>(true, null, successCom);
            }

            var failResponse =
                JsonConvert.DeserializeObject<FamisHttpResponse>(
                    await response.Content.ReadAsStringAsync());
            return new UpsertResponse<Company>(false, failResponse.Message, null);
        }

        public async Task<UpsertResponse<WorkOrder>> PostWo(CreateWorkOrder wo) {
            var body = JsonConvert.SerializeObject(wo);
            await AuthorizeIfNeeded();
            var url = _baseUri + "/apis/360facility/v1/workorders?";
            var content = new StringContent(body, Encoding.UTF8, "application/json");
            var response = await _client.PostAsync(url, content);

            if (response.IsSuccessStatusCode) {
                var successWo =
                    JsonConvert.DeserializeObject<WorkOrder>(
                        await response.Content.ReadAsStringAsync());
                return new UpsertResponse<WorkOrder>(true, null, successWo);
            }

            var failResponse =
                JsonConvert.DeserializeObject<FamisHttpResponse>(
                    await response.Content.ReadAsStringAsync());
            return new UpsertResponse<WorkOrder>(false, failResponse.Message, null);
        }

        public async Task<UpsertResponse<OtherCost>> PostOtherCost(OtherCost otherCost) {
            var body = JsonConvert.SerializeObject(
                otherCost, Formatting.None,
                new JsonSerializerSettings {
                    NullValueHandling = NullValueHandling.Ignore
                });
            await AuthorizeIfNeeded();
            var url = _baseUri + "/apis/360facility/v1/othercosts?";
            var content = new StringContent(body, Encoding.UTF8, "application/json");
            var response = await _client.PostAsync(url, content);

            if (response.IsSuccessStatusCode) {
                var successCost =
                    JsonConvert.DeserializeObject<OtherCost>(
                        await response.Content.ReadAsStringAsync());
                return new UpsertResponse<OtherCost>(true, null, successCost);
            }

            var failResponse =
                JsonConvert.DeserializeObject<FamisHttpResponse>(
                    await response.Content.ReadAsStringAsync());
            return new UpsertResponse<OtherCost>(false, failResponse.Message, null);
        }

        // method for persisting records to FAMIS
        public delegate Task<UpsertResponse<JsonMap>> PersistFunction(
            string url,
            StringContent content);

        public async Task<UpsertResponse<JsonMap>> CreateRecord(
            string endpoint,
            JsonMap obj) {
            var body = JsonConvert.SerializeObject(obj);
            await AuthorizeIfNeeded();
            var url = _baseUri + endpoint;
            var content = new StringContent(body, Encoding.UTF8, "application/json");
            return await UpsertRecord(url, content, createRecord());
        }

        public async Task<UpsertResponse<JsonMap>> UpdateRecord(
            string endpoint,
            JsonMap obj,
            int idValue) {
            var body = JsonConvert.SerializeObject(obj);
            await AuthorizeIfNeeded();
            // append id as the key for update
            var url = _baseUri + endpoint + $"?key={idValue}";
            var content = new StringContent(body, Encoding.UTF8, "application/json");
            return await UpsertRecord(url, content, updateRecord());
        }


        // This method gets the value form the given idField attribute
        // then removes it from the JsonMap because FAMIS will not allow it when creating 
        // or updating
        private static int getIdFromObj(JsonMap obj, string idField) {
            var id = (Int64) obj[idField];
            obj.Remove(idField);
            return Convert.ToInt32(id);
        }


        public Task<UpsertResponse<JsonMap>> UpsertRecord(
            string url,
            StringContent content,
            PersistFunction func) {
            return func(url, content);
        }

        private PersistFunction createRecord() {
            return async (url, content) => {
                var response = await _client.PostAsync(url, content);
                return await handleResponse(response);
            };
        }

        private PersistFunction updateRecord() {
            return async (url, content) => {
                var response = await _client.PatchAsync(url, content);
                return await handleResponse(response);
            };
        }

        private static async Task<UpsertResponse<JsonMap>> handleResponse(
            HttpResponseMessage response) {
            var responseBody = await response.Content.ReadAsStringAsync();
            if (response.IsSuccessStatusCode) {
                var record = JsonConvert.DeserializeObject<JsonMap>(responseBody);
                return new UpsertResponse<JsonMap>(true, null, record);
            }

            var failResponse = JsonConvert.DeserializeObject<FamisHttpResponse>(responseBody);
            return new UpsertResponse<JsonMap>(false, failResponse.Message, null);
        }
    }
}