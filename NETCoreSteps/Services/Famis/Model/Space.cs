using System;
using System.Collections.Generic;
using System.Text;

namespace Famis.Model
{

    public class Space
    {
        public int Id { get; set; }
        public string ExternalId { get; set; }
        public string Name { get; set; }
        public string FloorName { get; set; }
        public int? OccupantsDay { get; set; }
        public int? OccupantsNight { get; set; }
        public int? OccupantsTotal { get; set; }
        public decimal? Size { get; set; }
        public int? RoomTypeId { get; set; }
        public int? PropertyId { get; set; }
        public string PropertyExternalId { get; set; }
        public int? FloorId { get; set; }
        public string FloorExternalId { get; set; }
        public int? AccountId { get; set; }
        public int? ContactPersonId { get; set; }
        public string ContactPersonExternalId { get; set; }
        public int? CostCodeId { get; set; }
        public int? RateScheduleId { get; set; }
        public bool ActiveFlag { get; set; }
        public bool RequestCloseNotification { get; set; }
        public bool OverrideBillableFlag { get; set; }
        public int? TabOrder { get; set; }
        public bool SurveyFlag { get; set; }
        public bool CommonAreaFlag { get; set; }
        public int? OccupancyStatusId { get; set; }
        public DateTime UpdateDate { get; set; }
        public int UpdatedById { get; set; }
        public string CadId { get; set; }
        public int? ClassId { get; set; }
        public string Telephone { get; set; }
        public string CadSpaceId { get; set; }
        public string LongDescription { get; set; }
        public string EpPlanChildId { get; set; }
        public int? RequestPriorityListId { get; set; }
        
    }
}
