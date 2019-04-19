using System.Collections.Generic;
using System.Net.Http;
using System.Threading.Tasks;
using Famis.Model;
using JsonMap = System.Collections.Generic.Dictionary<string, object>;

namespace Famis
{
    public interface IService
    {
        Task<PagedResponse<Company>> GetCompanies(
            string filter = null,
            string expand = null,
            int offset = 0,
            int limit = -1);

        Task<Company> GetCompany(string id, string expand = null);

        Task<PagedResponse<Property>> GetProperties(
            string filter = null,
            string expand = null,
            int offset = 0,
            int limit = -1);

        Task<Property> GetProperty(string id, string expand = null);

        Task<PagedResponse<Floor>> GetFloors(string filter, string expand, int offset, int limit);
        Task<Floor> GetFloor(string id, string expand);

        Task<PagedResponse<Country>> GetCountries(
            string filter = null,
            string expand = null,
            int offset = 0,
            int limit = -1);

        Task<Country> GetCountry(string id, string expand = null);

        Task<PagedResponse<Crew>> GetCrews(
            string filter = null,
            string expand = null,
            int offset = 0,
            int limit = -1);

        Task<Crew> GetCrew(string id, string expand = null);

        Task<PagedResponse<InvoiceHeader>> GetInvoiceHeaders(
            string filter = null,
            string expand = null,
            int offset = 0,
            int limit = -1);

        Task<InvoiceHeader> GetInvoiceHeader(string id, string expand = null);

        Task<PagedResponse<InvoiceLine>> GetInvoiceLines(
            string filter = null,
            string expand = null,
            int offset = 0,
            int limit = -1);

        Task<InvoiceLine> GetInvoiceLine(string id, string expand = null);

        Task<PagedResponse<JournalEntry>> GetJournalEntries(
            string filter = null,
            string expand = null,
            int offset = 0,
            int limit = -1);

        Task<JournalEntry> GetJournalEntry(string id, string expand = null);

        Task<PagedResponse<MaterialItem>> GetMaterialItems(
            string filter = null,
            string expand = null,
            int offset = 0,
            int limit = -1);

        Task<MaterialItem> GetMaterialItem(string id, string expand = null);

        Task<PagedResponse<PoHeader>> GetPoHeaders(
            string filter = null,
            string expand = null,
            int offset = 0,
            int limit = -1);

        Task<PoHeader> GetPoHeader(string id, string expand = null);

        Task<PagedResponse<PoLine>> GetPoLines(string filter, string expand, int offset, int limit);
        Task<PoLine> GetPoLine(string id, string expand);

        Task<PagedResponse<Receipt>> GetPoReceipts(
            string filter = null,
            string expand = null,
            int offset = 0,
            int limit = -1);

        Task<Receipt> GetPoReceipt(string id, string expand = null);

        Task<PagedResponse<PoStatus>> GetPoStatuses(
            string filter = null,
            string expand = null,
            int offset = 0,
            int limit = -1);

        Task<PoStatus> GetPoStatus(string id, string expand = null);

        Task<PagedResponse<PrHeader>> GetPrHeaders(
            string filter = null,
            string expand = null,
            int offset = 0,
            int limit = -1);

        Task<PrHeader> GetPrHeader(string id, string expand = null);

        Task<PagedResponse<PrLine>> GetPrLines(
            string filter = null,
            string expand = null,
            int offset = 0,
            int limit = -1);

        Task<PrLine> GetPrLine(string id, string expand = null);

        Task<PagedResponse<State>> GetStates(
            string filter = null,
            string expand = null,
            int offset = 0,
            int limit = -1);

        Task<State> GetState(string id, string expand = null);

        Task<PagedResponse<Uom>> GetUoms(
            string filter = null,
            string expand = null,
            int offset = 0,
            int limit = -1);

        Task<Uom> GetUom(string id, string expand = null);

        Task<PagedResponse<AccountSegmentValueNPFA>> GetAcctSegs(
            string filter = null,
            string expand = null,
            int offset = 0,
            int limit = -1);

        Task<AccountSegmentValueNPFA> GetAcctSeg(string id, string expand = null);

        Task<PagedResponse<WorkOrder>> GetWorkOrders(
            string filter = null,
            string expand = null,
            int offset = 0,
            int limit = -1);

        Task<WorkOrder> GetWorkOrder(string id, string expand = null);

        Task<PagedResponse<Warehouse>> GetWarehouses(
            string filter = null,
            string expand = null,
            int offset = 0,
            int limit = -1);

        Task<PagedResponse<RequestType>> GetRequestTypes(
            string filter = null,
            string expand = null,
            int offset = 0,
            int limit = -1);

        Task<PagedResponse<OtherCostType>> GetOtherCostTypes(
            string filter = null,
            string expand = null,
            int offset = 0,
            int limit = -1);

        Task<Warehouse> GetWarehouse(string id, string expand = null);

        Task<PagedResponse<User>> GetUsers(
            string filter = null,
            string expand = null,
            int offset = 0,
            int limit = -1);

        Task<User> GetUser(string id, string expand = null);

        Task<UpsertResponse<User>> PostUser(User user);

        Task<UpsertResponse<User>> PatchUser(User user, int id);


        Task<PagedResponse<Receipt>> GetReceipts(
            string filter = null,
            string expand = null,
            int offset = 0,
            int limit = -1);

        Task<PagedResponse<RequestPriority>> GetRequestPriorities(
            string filter = null,
            string expand = null,
            int offset = 0,
            int limit = -1);

        Task<PagedResponse<CrewUserAssociation>> GetCrewUserAssociations(
            string filter = null,
            string expand = null,
            int offset = 0,
            int limit = -1);

        Task<PagedResponse<LaborEntry>> GetLaborEntries(
            string filter = null,
            string expand = null,
            int offset = 0,
            int limit = -1);

        Task<PagedResponse<LaborReason>> GetLaborReasons(
            string filter = null,
            string expand = null,
            int offset = 0,
            int limit = -1);

        Task<PagedResponse<AccountSegmentNPFA>> GetAcctSegNpfas(
            string filter = null,
            string expand = null,
            int offset = 0,
            int limit = -1);

        Task<AccountSegmentNPFA> GetAcctSegNpfa(string id, string expand = null);

        Task<Receipt> GetReceipt(string id, string expand = null);

        Task<PagedResponse<T>> GetNext<T>(PagedResponse<T> current);

        Task<UpsertResponse<InvoiceHeader>> PostInvoice(InvoiceHeader famisInv);

        Task<UpsertResponse<InvoiceLine>> PostInvoiceLine(InvoiceLine line);

        Task<UpsertResponse<PoHeader>> PostPo(PoHeader famisPo);

        Task<UpsertResponse<PoLine>> PostPoLine(POLineCreateRequest line);

        Task<UpsertResponse<WorkOrder>> PatchProjectWriteBack(string json, int Id);

        Task<UpsertResponse<WorkOrder>> PostProjectCW(ChildWorkOrder wo);

        Task<UpsertResponse<OtherCost>> PostOtherCost(OtherCost otherCost);

        Task<UpsertResponse<Company>> PatchCompany(Company company, int Id);

        Task<UpsertResponse<Company>> PostCompany(Company com);

        Task<UpsertResponse<WorkOrder>> PostWo(CreateWorkOrder wo);

        Task<UpsertResponse<JsonMap>> CreateRecord(string endpoint, JsonMap obj);
        Task<UpsertResponse<JsonMap>> UpdateRecord(string endpoint, JsonMap obj, int id);
    }
}