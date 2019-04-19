using Newtonsoft.Json;
using System;
using System.Collections.Generic;
using System.ComponentModel.DataAnnotations.Schema;
using System.Linq;
using System.Runtime.Serialization;
using System.Web;

namespace Famis.Model {

    public class LaborEntry {
        public int Id { get; set; }
        public double? TotalHours { get; set; }
        public string TimeType { get; set; }
        public int? ActivityId { get; set; }
        public string ActivityExternalId { get; set; }
        public string ActivityName { get; set; }
        public int? PropertyId { get; set; }
        public string PropertyExternalId { get; set; }
        public string PropertyName { get; set; }
        public string Comments { get; set; }
        public int? RequestId { get; set; }
        public string RequestExternalId { get; set; }
        public int? UserId { get; set; }
        public string UserExternalId { get; set; }
        public string EntryDate { get; set; }
        public int? PayPeriodId { get; set; }
        public int? PayYear { get; set; }
        public int? CrewId { get; set; }
        public string CrewExternalId { get; set; }
        public string StartTime { get; set; }
        public string EndTime { get; set; }
        public int? PositionId { get; set; }
        public string PositionExternalId { get; set; }
        public int? LaborReasonId { get; set; }
        public int? StatusId { get; set; }
        public LaborEntryStatus Status { get; set; }
    }

    public class LaborEntryStatus {
        public int StatusId { get; set; }
        public string Name { get; set; }
        public bool ReadyFlag { get; set; }
        public bool PendingFlag { get; set; }
        public bool ApprovedFlag { get; set; }
        public bool RejectedFlag { get; set; }
    }

}